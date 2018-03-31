
构建包括扩展的Fn Server镜像：
    
    vim ext.yaml

```yaml
extensions:
  - name: github.com/iwangxiaodong/container-config/fn/myext
```
  
    fn build-server -t cr.openle.com/my/fn-server-with-ext:v1.0.0
    
    docker push cr.openle.com/my/fn-server-with-ext:v1.0.0
    
    kubectl set image deployment/my-fns my-fns-p=cr.openle.com/my/fn-server-with-ext:v1.0.0
