
构建包括扩展的Fn Server镜像：

    curl -LSs https://raw.githubusercontent.com/fnproject/cli/master/install | sh

    mkdir my-fn-ext; cd my-fn-ext
    
    curl https://raw.githubusercontent.com/iwangxiaodong/container-config/master/fn/myext/ext.yaml
   
    fn build-server -t gcr.io/$GOOGLE_CLOUD_PROJECT/fn-server-with-ext:1.0.0
    
    gcloud auth configure-docker
        
    docker push gcr.io/$GOOGLE_CLOUD_PROJECT/fn-server-with-ext:1.0.0
    
    kubectl set image deployment/my-fns my-fns-p=gcr.io/$GOOGLE_CLOUD_PROJECT/fn-server-with-ext:1.0.0
