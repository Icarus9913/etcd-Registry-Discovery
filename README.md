## ETCD V3 服务注册与发现  

discovery包中的服务为service,监听为master:    
    * service启动时向etcd注册自己的信息,注册到services/  这个目录      
    * service可能异常退出,需要维护一个TTL(V3使用lease实现), 类似心跳, 如果挂了master可以监听到    
    * master监听services/ 目录下的所有服务,根据不同类型的action(如put/delete),进行处理  

service注册:  
    * 提供key(service name), value(serviceinfo)进行registry      
    * start后,执行keeplive, 挂掉或者stop后执行revoke.(相当于unregistry)      

master监听:  
    * 提供监听路径path,启动master,当put时加入map,delete时从map中删除.    
