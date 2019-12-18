-- Part 1
import System.IO (isEOF)
import Data.List.Split
----------------------------------

data Box = Box Int Int Int deriving (Show)

area :: Box -> Int
area (Box 0 0 0)      = 0
area (Box l w h) = (2*l*w) + (2*w*h) + (2*h*l)

splitString :: String -> [String]
splitString s = splitOn "x" s


-- Main
main :: IO()
main = readIn

readIn = do
  done <- isEOF
  if done
    then putStrLn ""
    else do inp <- getLine
            print $ splitString inp
            readIn


