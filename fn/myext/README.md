
构建包括扩展的Fn Server镜像：

    curl -LSs https://raw.githubusercontent.com/fnproject/cli/master/install | sh

    mkdir my-fn-ext; cd my-fn-ext
    
    wget https://raw.githubusercontent.com/iwangxiaodong/container-config/master/fn/myext/ext.yaml
    
    注意：由于是从源码构建，so只在官网构建状态为“成功”时构建
   
    sudo fn --verbose build-server -t gcr.io/$GOOGLE_CLOUD_PROJECT/fn-server-with-ext:1.0.0-b1
    
    sudo su root
    
    gcloud auth configure-docker
        
    docker push gcr.io/$GOOGLE_CLOUD_PROJECT/fn-server-with-ext:1.0.0-b1
    
    exit
    
    kubectl set image deployment/my-fns fns=gcr.io/$GOOGLE_CLOUD_PROJECT/fn-server-with-ext:1.0.0-b1 --namespace=all
