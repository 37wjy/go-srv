sudo yum install -y yum-utils
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
sudo yum install --allowerasing docker-ce docker-ce-cli containerd.io


# (尝试在命令行中添加 '--allowerasing' 来替换冲突的软件包 或 '--skip-broken' 来跳过无法安装的软件包 或 '--nobest' 来不只使用软件包的最佳候选)