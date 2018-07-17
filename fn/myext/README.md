
构建包括扩展的Fn Server镜像：

    curl -LSs https://raw.githubusercontent.com/fnproject/cli/master/install | sh

    mkdir my-fn-ext; cd my-fn-ext
    
    wget https://raw.githubusercontent.com/iwangxiaodong/container-config/master/fn/myext/ext.yaml
    
    注意：由于是通过dep或go get分析import package获取master源码，so只在官网构建状态为“PASSED”时创建Server
   
    sudo fn --verbose build-server -t gcr.io/$GOOGLE_CLOUD_PROJECT/fn-server-with-ext:1.0.0-b1
    
    sudo su root
    
    gcloud auth configure-docker
        
    docker push gcr.io/$GOOGLE_CLOUD_PROJECT/fn-server-with-ext:1.0.0-b1
    
    exit
    
    kubectl set image deployment/my-fns fns=gcr.io/$GOOGLE_CLOUD_PROJECT/fn-server-with-ext:1.0.0-b1 --namespace=all


【例外情况】：

    若生成的tmp/main.go与fn项目不一致则以后者为准：
    
        https://raw.githubusercontent.com/fnproject/fn/master/cmd/fnserver/main.go
        
        删除重复导入：	_ "golang.org/x/net/trace"
        
                import _ "github.com/iwangxiaodong/container-config/fn/myext"
        
        [在funcServer.Start(ctx)上方增加]：
        
                funcServer.AddExtensionByName("github.com/iwangxiaodong/container-config/fn/myext")
        
        
    sudo docker build -t gcr.io/$GOOGLE_CLOUD_PROJECT/fn-server-with-ext:1.0.0-b1 .
    
    sudo gcloud auth configure-docker
    
    docker push gcr.io/$DEVSHELL_PROJECT_ID/fn-server-with-ext:1.0.0-b1
