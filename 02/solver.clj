(ns solve
  (:require [clojure.string :as str]))

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
