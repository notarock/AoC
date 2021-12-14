(ns solve
  (:require [clojure.string :as str]))

;; Part 1

(defn count-occurence [word letter]
  (->> word
       (re-seq (re-pattern letter))
       count))

(defn count-valid [min max letter password]
  (let [occurence (count-occurence password letter)]
    (if (and (<= occurence max) (>= occurence min))
      1
      0)))

(defn validate-password [password-line]
  (let* [password (str/split password-line #" ")
         min (-> password (first) (str/split #"-") (first) (read-string))
         max (-> password (first) (str/split #"-") (second) (str/join) (read-string))
         letter (-> password (second) (first) (str))
         word (-> password (last) (str))]
    (count-valid min max letter word)))

(let [input (str/split (slurp "input") #"\n")]
  (reduce + (map validate-password input)))

;; Part 2

(defn validate-second-password [password-line]
  (let* [password (str/split password-line #" ")
         pos1 (-> password (first) (str/split #"-") (first) (read-string))
         pos2 (-> password (first) (str/split #"-") (second) (str/join) (read-string))
         letter (-> password (second) (first) (str))
         word (-> password (last) (str))]
    (bit-xor (position-valid pos1 letter word) (position-valid pos2 letter word))))

(defn position-valid [pos letter password]
  (if (= (str (nth password (- pos 1))) letter)
    1
    0))

(let [input (str/split (slurp "input") #"\n")]
  (reduce + (map validate-second-password input)))
