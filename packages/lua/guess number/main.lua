math.randomseed(os.time())
number = math.random(1, 100)

player = {} -- means player is the table
player.guess = 0

while player.guess ~= number do
    print("Enter number from 1 to 100")
    player.answer = io.read()
    player.guess = tonumber(player.answer)

    if player.guess > number then
        print("Too High")
    elseif player.guess < number then
        print("Too Low")
    else 
        print("You guessed it!!")
    end
end

print(number) --logs the answer