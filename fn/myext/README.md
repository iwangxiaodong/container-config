
构建包括扩展的Fn Server镜像：

    curl -LSs https://raw.githubusercontent.com/fnproject/cli/master/install | sh

    mkdir my-fn-ext; cd my-fn-ext
    
    wget https://raw.githubusercontent.com/iwangxiaodong/container-config/master/fn/myext/ext.yaml
   
    sudo fn --verbose build-server -t gcr.io/$GOOGLE_CLOUD_PROJECT/fn-server-with-ext:1.0.0-b1
    
    sudo su root
    
    gcloud auth configure-docker
        
    docker push gcr.io/$GOOGLE_CLOUD_PROJECT/fn-server-with-ext:1.0.0-b1
    
    kubectl set image deployment/my-fns fns=gcr.io/$GOOGLE_CLOUD_PROJECT/fn-server-with-ext:1.0.0-b1 --namespace=all
