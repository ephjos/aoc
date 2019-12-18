-- Part 1

getFloor :: [Char] -> Int -> Int
getFloor [] count = count
getFloor (x:xs) count
  | x == '(' = getFloor xs count+1
  | x == ')' = getFloor xs count-1
  | otherwise = 0

main :: IO()
main = do
    name <- getLine
    print (getFloor name 0)
