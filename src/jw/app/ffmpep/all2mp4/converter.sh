#!/usr/bin

input_dir=$1
output_dir=$2

if [ $# != 2 ]; then
   echo "please input 2 parameters, 1st for input folder, 2nd for output folder"
   exit 1
fi

for item in `find ${input_dir} -maxdepth 1 -type f -name "*.mkv" -exec basename -s .mkv {} \;`
do
	#echo "item => ${item}"
	#name=${item:18}".mp4"
	echo "find . -maxdepth 1 -type f -name \"*${item}.mkv\" -exec ffmpeg -i {} -vcodec copy /f/tmp/${item}.mp4 \;"
	find ${input_dir} -maxdepth 1 -type f -name "*${item}.mkv" -exec ffmpeg -i {} -vcodec copy ${output_dir}/${item}.mp4 \;
	echo "ret: "$?

done

