package main

import (
	"context"
	"github.com/davyxu/golog"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	Lg        = GetLogger()
	clientset *kubernetes.Clientset
)

func GetLogger() *golog.Logger {
	lg := golog.New("")
	lg.EnableColor(true)
	lg.SetParts(golog.LogPart_Level, golog.LogPart_TimeMS, golog.LogPart_ShortFileName)
	return lg
}

func k8sInit() {
	kuberconfig := "xxx"
	config, err := clientcmd.BuildConfigFromFlags("", kuberconfig)
	if err != nil {
		config, err = clientcmd.BuildConfigFromFlags("", "")
		if err != nil {
			Lg.Errorf(err.Error())
			return
		}
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		Lg.Errorf(err.Error())
		return
	}

}

func GetTargetPod(ns, podName string) (pod *v1.Pod, err error) {
	pod, err = clientset.CoreV1().Pods(ns).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		Lg.Errorf(err.Error())
		return nil, err
	}
	return pod, nil
}

func PatchRSorRC(ns, podName string) error {
	pod, err := GetTargetPod(ns, podName)
	if err != nil {
		Lg.Errorf(err.Error())
		return err
	}
	switch pod.GetObjectMeta().GetOwnerReferences()[0].Kind {
	case "ReplicaSet":
		rs, err := clientset.ExtensionsV1beta1().ReplicaSets(ns).Get(context.TODO(), pod.GetObjectMeta().GetOwnerReferences()[0].Name, metav1.GetOptions{})
		if err != nil {
			Lg.Errorf(err.Error())
			return err
		}
		payloadBytes := []byte(`{"spec": {"template": {"spec": {"nodeSelector": {"node_pool": "dev"}}}}}`)
		_, err = clientset.AppsV1().Deployments(ns).Patch(
			context.TODO(),
			rs.GetOwnerReferences()[0].Name,
			types.JSONPatchType,
			payloadBytes,
			metav1.PatchOptions{},
		)
		if err != nil {
			Lg.Errorf(err.Error())
			return err
		} else {
			Lg.Infof("rs patched done")
		}
		rs = nil
	case "ReplicationController":

	default:
		Lg.Infof("default: %s", pod.GetObjectMeta().GetOwnerReferences()[0].Name)
	}
	pod = nil
	return nil
}

func main() {
	k8sInit()
	PatchRSorRC("", "")

}
