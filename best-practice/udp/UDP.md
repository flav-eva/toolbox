

必须清楚知道你的UDP是连接的(connected)还是未连接(unconnected)的，这样你才能正确的选择的读写方法。

* 如果*UDPConn是connected,读写方法是Read和Write。
* 如果*UDPConn是unconnected,读写方法是ReadFromUDP和WriteToUDP（以及ReadFrom和WriteTo)



参考：

https://colobu.com/2016/10/19/Go-UDP-Programming/
