package interview

//机器人坐标问题
//
//问题描述
//有一个机器人,给一串指令,L左转 R右转,F前进一步,B后退一步,问最后机器人的坐标,
//
//最开始,机器人位于 0 0,方向为正Y. 可以输入重复指令n .比如 R2(LF) 这个等于指令 RLFLF. 问最后机器人的坐标是多少.
//解题思路
//这里的一个难点是解析重复指令.主要指令解析成功,计算坐标就简单了.
