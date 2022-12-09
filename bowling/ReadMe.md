#The Bowling Game
_An exercise in Test Driven Development with Go_

Based on the demonstration by Robert "Uncle Bob" Martin: http://www.butunclebob.com/ArticleS.UncleBob.TheBowlingGameKata


##Scoring Bowling
```
 Score for first ball in frame
 |
 |    Score for second ball in frame (SP = Spare)
 |    |
 v    V
|1    4|4    5|6   SP|5   SP|STRIKE|0    1|7   SP|6   SP|STRIKE|2  SP6|
|   5  |  14  |  29  |  49  |  60  |  61  |  77  |  97  |  117 |  133 |  <-- Running total
```

A game of Ten Pin Bowling consists of 10 **frames** as shown above.  In each frame the 
player has two opportunities to knock down 10 **pins**.  The score for the frame is 
the total number of pins knocked down, plus bonuses for **strikes** and **spares**.

A **spare** is when the player knocks down all 10 pins in two tries.  The bonus for
that frame is the number of pins knocked down by the next roll.  So in frame 3
above, the score is 10 (the total number knocked down) plus a bonus of 5 (the
number of pins knocked down on the next roll.)

A **strike** is when the player knocks down all 10 pins on his first try.  The bonus
for that frame is the value of the next two balls rolled.

In the tenth frame a player who rolls a spare or strike is allowed to roll the extra
balls to complete the frame.  However no more than three balls can be rolled in
tenth frame.

###Some other terms:
* **Gutter game**: A game with zero score
* **Perfect game**: A game with the highest possible score of 300.

##The Requirements
Write a class named `Game` that has two methods:
* `Roll(pins int)` is called each time the player rolls a ball.  The argument is the number of pins knocked down.
* `Score() int` is called only at the very end of the game.  It returns the total score for that game.

##Begin
* Create a project named BowlingGame
* Create a unit test named BowlingGameTest
