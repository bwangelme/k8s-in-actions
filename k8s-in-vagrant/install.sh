#!/bin/bash
#
# Author: bwangel<bwangel.me@gmail.com>
# Date: 5,24,2020 16:38

echo "设置 apt"
cp /vagrant/conf/sources.list /etc/apt/sources.list
# 设置 apt 安装的代理
cat >> /etc/apt/apt.conf.d/proxy.conf <<EOF
Acquire {
  HTTP::proxy "http://10.8.0.1:8118";
  HTTPS::proxy "http://10.8.0.1:8118";
}
EOF

echo '设置主机名的解析'
cat >> /etc/hosts <<EOF
172.17.9.101 k8s-node1
172.17.9.102 k8s-node2
172.17.9.103 k8s-node3
EOF

echo '安装Docker'
export DEBIAN_FRONTEND=noninteractive
# step 1: 安装必要的一些系统工具
apt-get update && apt-get install -y apt-transport-https ca-certificates curl software-properties-common gnupg2
# step 2: 安装GPG证书
curl -fsSL http://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
# Step 3: 写入软件源信息
add-apt-repository "deb [arch=amd64] http://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"
# Step 4: 更新并指定版本安装Docker-CE
apt-get update && apt-get install -y \
      containerd.io=1.2.13-1 \
      docker-ce=5:19.03.8~3-0~ubuntu-$(lsb_release -cs) \
      docker-ce-cli=5:19.03.8~3-0~ubuntu-$(lsb_release -cs)
# Set up the Docker daemon
# TODO: master 是 systemd，从节点是 cgroupfs
cat > /etc/docker/daemon.json <<EOF
{
  "exec-opts": ["native.cgroupdriver=cgroupfs"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2"
}
EOF
# 设置 Docker 的代理
[[ ! -d "/etc/systemd/system/docker.service.d" ]] && mkdir -p /etc/systemd/system/docker.service.d
cat > /etc/systemd/system/docker.service.d/http-proxy.conf <<EOF
[Service]
Environment="HTTP_PROXY=10.8.0.1:8118"
Environment="HTTPS_PROXY=10.8.0.1:8118"
Environment="NO_PROXY=localhost,127.0.0.0/8,10.0.0.0/8,172.17.0.0/16,192.168.0.0/16"
EOF
# 创建 Docker 用户组
egrep "^docker" /etc/group >& /dev/null
if [ $? -ne 0 ]
then
  groupadd docker
fi
usermod -aG docker vagrant

echo '关闭 Swap'
swapoff -a && sed -i '/ swap / s/^/#/' /etc/fstab


echo '安装 Kubernetes'
export HTTP_PROXY='10.8.0.1:8118' HTTPS_PROXY='10.8.0.1:8118'
export NO_PROXY=localhost,127.0.0.0/8,10.0.0.0/8,172.17.0.0/16,192.168.0.0/16
curl -fsSL https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
unset HTTP_PROXY HTTPS_PROXY
cat >/etc/apt/sources.list.d/kubernetes.list <<EOF
deb https://packages.cloud.google.com/apt/ kubernetes-xenial main
EOF
apt-get update && apt-get install -y kubelet kubeadm kubectl
apt-mark hold kubelet kubeadm kubectl

echo '启动k8s'
if [[ $1 == 1 ]];then
    kubeadm init --apiserver-advertise-address 172.17.9.101 --pod-network-cidr=10.244.0.0/16
    KUBECONFIG=/etc/kubernetes/admin.conf kubectl apply -f /vagrant/k8s-svc/kube-flannel.yml
fi

# TODO: join 需要 k8s 的 Token，这个如何设置

# Your Kubernetes control-plane has initialized successfully!

# To start using your cluster, you need to run the following as a regular user:

#   mkdir -p $HOME/.kube
#   sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
#   sudo chown $(id -u):$(id -g) $HOME/.kube/config

# You should now deploy a pod network to the cluster.
# Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
#   https://kubernetes.io/docs/concepts/cluster-administration/addons/

# Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 172.17.9.101:6443 --token 8p1l90.rpdpufdstui9nwvf \
    --discovery-token-ca-cert-hash sha256:e32c36db0b8ce0cbb88ba6016784f4d0d5e7851978de6f150e84f9337c37b28a
