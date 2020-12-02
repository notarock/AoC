(ns aoc
  (:require [clojure.string :as str]))

(let [input (str/split (slurp "input") #"\n")]
  (loop [x (- (count input) 1)]
    (when (and (not= x (count input)) (not= x 0))
      (loop [y (- (count input) x 1)]
        (let [current (read-string (nth input x))
              second (read-string (nth input y))]
          (if (= (+ current second) 2020)
            (println (* current second))
            (when (not= y 0)
              (recur (dec y))))))
      (recur (dec x)))))
