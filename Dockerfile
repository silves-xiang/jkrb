FROM centos:latest
ADD ./main /main
ENV TZ=Asia/Shanghai
ENTRYPOINT [ "./main" ]