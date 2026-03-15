-- https://luabyexample.techplexlabs.com

-- print('hellow world')

-- print("conca".."tinated")

-- print("1+1=", 1+1)
-- print("7.0/6.0=", 7.0/6.0)

-- print(true or false)
-- print(true and false)
-- print(not false)

-- print(true or (true and false))

-- if nil then
--     print("if is nil")
-- else 
--     print('else if not nil')
-- end

-- -- variables

-- str = "some string"
-- a = 1
-- f = 2.0

-- print(str)
-- print(a)
-- print(f..a)

-- local temp = "temp value"
-- print("temp var doesnt pollute the global scopes",temp)

-- local nilll = nil
-- print(nilll)

-- for loop

-- for i=1, 10, 1 do 
--     print(i)
-- end

-- -- i is local to for loop 
-- print(i) -- logs nil

-- -- break

-- for i=1, 100, 2 do
--     if i == 5 then
--         break
--     end
--     print(i)
-- end

-- --  lua doesnt have continue (workaround is goto)
-- -- print only odd numbers
-- for i=1, 10, 1 do
--     if i % 2 == 0 then goto continue end
--     print(i)
--     ::continue::
-- end

-- -- while loop

-- local i = 1

-- while i < 10 do 
--     print(i)
--     i = i + 1
-- end

-- while true do 
--     print("this loop will run continuously")
--     break
-- end

-- if else

-- local no = 9

-- if no < 0 then
--     print("negative number")
-- elseif no < 10 then
--     print("has 1 digit")
-- else 
--     print("positive number")
-- end

-- if no % 2 == 0 then
--     print("even number")
-- else 
--     print("odd number")
-- end

-- functions

-- function add(a, b)
--     return a + b
-- end

-- print(add(1,2))

-- multiple values

-- arr = {'a', 'b', 'c'}
-- for i,v in ipairs(arr) do
--     print("i:"..i, "v:"..v)
-- end
-- function number()
--     return 1, 3
-- end

-- a, b = number()

-- print(a, b)

-- _, b = number()

-- print(b)

-- -- variadic function
-- local function sum(...)
--     print("items:", ...)

--     total = 0 
--     for i, v in ipairs({...}) do 
--         total = total + v
--     end
--     print("sum:", total)
-- end

-- sum(1,2)
-- sum(1,2,3,4,5)

