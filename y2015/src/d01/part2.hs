-- Part 2

getFloor :: [Char] -> Int -> Int -> Int
getFloor [] floor pos = pos
getFloor (x:xs) floor pos
  | floor < 0    = pos
  | x == '('     = getFloor xs (floor+1) (pos+1)
  | x == ')'     = getFloor xs (floor-1) (pos+1)
  | otherwise    = 0

main :: IO()
main = do
    name <- getLine
    print (getFloor name 0 0)
