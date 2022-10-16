# API

這些 api 都已經包含用戶資訊。

> 喊數字  
/api/call/:number
> 將[?色機器人]往[方向]移動  
/api/move/:color/:direction
> 發起挑戰  
/api/challenge  

遊戲大概會這樣進行

clientA /api/call/7
clientB /api/call/7
clientC /api/call/6
clientA /api/call/5
==== time's up ====
clientA /api/move/:color/:direction * n
=== system judge ==
if success {
    while someone /api/challenge {
        if challenge success {
            challenger is the current winner
        } 
    }
    the current winner wins
} else {
    next client's move
}
