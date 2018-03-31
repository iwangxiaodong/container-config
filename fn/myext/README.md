
构建包括扩展的Fn Server镜像：
    
    vim ext.yaml

```yaml
extensions:
  - name: github.com/iwangxiaodong/container-config/fn/myext
```
  
    fn build-server -t cr.openle.com/my/fn-server-with-ext
