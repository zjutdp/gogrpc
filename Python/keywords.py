import keyword

str = "There are {} keywords in Python! They are {} ."
words = keyword.kwlist

print str.format(len(words), words)
