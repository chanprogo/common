



# 方法一

1. Ctrl + V 进入VISUAL BLOCK 模式，
2. 按 j （向下选取列）或者 k （向上选取列），

* 批量删除字符：直接（不用进入编辑模式）按 x 或者 d 就可以直接删去，再按 Esc 退出。
* 批量插入字符：再按 Shift + i 进入编辑模式然后输入你想要插入的字符（任意字符），再按两次 Esc 就可以实现批量插入字符。



# 方法二  

批量插入字符： 命令行模式下，输入 `:首行号,尾行号s/^/字符/g`  
批量删除字符： 命令行模式下，输入 `:首行号,尾行号s/^字符//g`  