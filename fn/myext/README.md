
构建包括扩展的Fn Server镜像：

    wget https://raw.githubusercontent.com/iwangxiaodong/container-config/master/fn/myext/myext.go    
    vim ext.yaml
extensions:
  - name: github.com/iwangxiaodong/container-config/fn/my-ext
  
    fn build-server -t cr.openle.com/my/fn-server-with-ext