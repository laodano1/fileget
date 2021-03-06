package main

import (
	"context"
	"fileget/util"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"net"
)

const (
//	rootCAFile = `
//-----BEGIN CERTIFICATE-----
//MIICyDCCAbCgAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl
//cm5ldGVzMB4XDTIwMDIyODExMDIwN1oXDTMwMDIyNTExMDIwN1owFTETMBEGA1UE
//AxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANbR
//Yrgv4abPYPVS3Pax7PRdMMyrpLSJ7f/F+2qCJHnVKTHeBTBa+pzIJcPHz0IOfyKC
//zbht9slPR3cJ3S/ebuNFeMD8GitqZ4RB0wLFGYKw3yCKvQcsAZNcooZK1CbVRF7z
//nsP0M1p71CCdzjKK5TCbaphYO6w3/XQM/v8CnliiXgxCFOuZl4O0AUaBDVJwz7xt
//J+qakQQdQW41FMDs8X6YDhY1BOPopTLdt2O7lhVnD8vBw1tP2NPgBp5WOWE8ZlMc
//mwyZpFtZDkNOhqH1h+xIsFfiwJhFnlj3qNSySoAnDDqLsdhJBNJ/H9T8OB3yi1WJ
//8z6Kt3YXx2p8S8awcgcCAwEAAaMjMCEwDgYDVR0PAQH/BAQDAgKkMA8GA1UdEwEB
///wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAJqes8oeb3nzwuREhyC1DkKV4BlD
//KXuqolgCaglFY/11VZp0exR30qmBsjpS0cCWLTchjOwc468uAumCNHPNT4XrEX0c
//PMSSZE6Gyvru1eZE6+yniPFVyq9cm9Rr74qaKwef1iPvmem0QNDekaeNsDWIJQ3g
//uUh9OmDE5dgpUq0d9H8ogb+sE0Ftry1O50FKvraR8zPgVzpnFU/Xmsk3lUExJ8rJ
//DreKWcUBEcVxKf1J9ao16ckb5Hk2RHzw7GUARyGB/o2chW3m4/+qKxIKy/Pv1eQ7
//ZAq8YMQFmEcLmXmgtdVSZeR5q0NKimJpL6+gooI2OSA4sGr3wvJTWO6TIo0=
//-----END CERTIFICATE-----
//`
//	tokenFile = `eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJ3YW5kYS1jZS1odWEiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlY3JldC5uYW1lIjoiZGVmYXVsdC10b2tlbi1qeHQyaiIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiNWMyZDY5YTMtNjJkZS00YmFkLTk5ODYtMmE3MGFmNGEyNzUzIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OndhbmRhLWNlLWh1YTpkZWZhdWx0In0.PljRRmyy6Pzhvue5l0wOO2NuQJg-rp1Hc1Tz89IH9ieN5PnOoba3VZgwnp1O6tX8NWeqyxnOxfYR0V4ihGkXrW0DMZsLec-9i_NrE4gGVzb0AHKPPm7upgzyvCV0yT3_SFcPFEAfjqx14_4ru6j0aBGc5aoCwGy4IkupkSwLRvkGrmS34cWm4qLp1595LY4N_Q0SIEIHEIUOmEzKcphhuBeWOqHb3JrCY0BwDZhNG9Y5eHlJ_oTNv9UflAW95IttYrUKf3PkSi-SNRSM7zHHRAVUMvJcUgnmgIVRwiXA7_r5yGwM6FpOnneJpNNrcSFSL6E0q-YK65XVb9ZH8ut0Wg`

	namespace = "wanda-games"
)


func Getk8sCliConfig() (*rest.Config, error) {
	host, port := "10.0.0.200", "6443"
	const (
		//tokenFile  = "/var/run/secrets/kubernetes.io/serviceaccount/token"
		tokenFile  = "./token"
		//rootCAFile = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
		rootCAFile = "./ca.crt"
	)

	//token, err := ioutil.ReadFile(tokenFile)
	//if err != nil {
	//
	//	return nil, err
	//}

	tlsClientConfig := rest.TLSClientConfig{}

	//if _, err := certutil.NewPool(rootCAFile); err != nil {
	//	util.Lg.Errorf("Expected to load root CA config from %s, but got err: %v", rootCAFile, err)
	//} else {
	//	tlsClientConfig.CAFile = rootCAFile
	//}
	tlsClientConfig.CAFile = rootCAFile

	return &rest.Config{
		Host:            "https://" + net.JoinHostPort(host, port),
		TLSClientConfig: tlsClientConfig,
		//BearerToken:     string(token),
		BearerToken:     tokenFile,
		BearerTokenFile: tokenFile,
	}, nil

}

func int32Ptr(i int32) *int32 { return &i }

func main() {
	//config, err := rest.InClusterConfig()
	config, err := Getk8sCliConfig()
	if err != nil {
		util.Lg.Errorf("Getk8sCliConfig error:", err.Error())
		return
	}


	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		util.Lg.Errorf("NewForConfig error: %s", err)
		return
	}

	pods, err := clientset.CoreV1().Pods(namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		util.Lg.Errorf("get Pods list error: %s", err)
		return
	}

	for _, pod := range pods.Items {
		util.Lg.Infof(" pod: %-19v, Image: %-44v, create time: %v", pod.Name, pod.Spec.Containers[0].Image, pod.GetCreationTimestamp())
	}


	svcs, err := clientset.CoreV1().Services(namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		util.Lg.Errorf("get Services list error: %s", err)
		return
	}

	for _, svc := range svcs.Items {
		util.Lg.Infof("service: %v", svc.Name)
	}

	sects, err := clientset.CoreV1().Secrets(namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		util.Lg.Errorf("get Secrets list error: %s", err)
		return
	}

	for _, scr := range sects.Items {
		util.Lg.Infof("Secret: %v", scr.Name)
	}

	// service account permission denied for accessing node info
	//nodes, err := clientset.CoreV1().Nodes().List(context.Background(), v1.ListOptions{})
	//if err != nil {
	//	util.Lg.Errorf("get nodes list error: %s", err)
	//	return
	//}
	//
	//for _, node := range nodes.Items {
	//	util.Lg.Infof("node: %v", node.Name)
	//}

	statefulset := &appsv1.StatefulSet{
		ObjectMeta: v1.ObjectMeta{
			Name: "my-alpine",
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: int32Ptr(2),
			Selector: &v1.LabelSelector{
				MatchLabels: map[string]string{
					"jw.app.name": "my-alpine",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Labels: map[string]string{
						"jw.app.name": "my-alpine",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name: "jw.alpine",
							Image: "alpine:latest",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "mycontainerport",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 7878,
								},
							},
						},
					},
				},
			},
		},
	}

	creResult, err := clientset.AppsV1().StatefulSets(namespace).Create(context.Background(), statefulset, v1.CreateOptions{})
	if err != nil {
		util.Lg.Errorf("create StatefulSet error: %s", err)
		return
	}

	util.Lg.Debugf("create stateful set time: %v", creResult.GetObjectMeta().GetCreationTimestamp())

	stss, err := clientset.AppsV1().StatefulSets("wanda-games").List(context.Background(), v1.ListOptions{})
	if err != nil {
		util.Lg.Errorf("get StatefulSets list error: %s", err)
		return
	}

	for _, sts := range stss.Items {
		util.Lg.Infof("StatefulSet: %v", sts.Name)
	}

	//wathcher, err := clientset.AppsV1().StatefulSets("wanda-games").Watch(context.Background(), v1.ListOptions{})
	//if err != nil {
	//	util.Lg.Errorf("start AppsV1 StatefulSets watch error: %s", err)
	//	return
	//}

	//tkDone := time.Tick(2 * time.Minute)
	//tk := time.Tick(5 * time.Second)
	//
	//wt := func() {
	//	for {
	//		select {
	//		case evt := <- wathcher.ResultChan():
	//			util.Lg.Debugf("objects: %v", evt.Object.GetObjectKind().GroupVersionKind())
	//		case <- tk:
	//		    util.Lg.Debugf("this is heartbeat!")
	//		case <- tkDone:
	//			util.Lg.Debugf("ticker timeout! jump out for loop")
	//			return
	//		}
	//	}
	//}

	//wt()

	util.Lg.Debugf("bye bye!")

}
