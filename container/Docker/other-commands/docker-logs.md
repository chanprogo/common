
`docker logs [OPTIONS] CONTAINER`  

```
--follow , -f		Follow log output
```

When you're done watching the logs, exit out by hitting `Ctrl+C`.  



`docker logs c90 --tail 9000 2>&1 | grep cipherText`  
