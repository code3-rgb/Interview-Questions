def getLongest(s, k):
    high = 0
    windows = set()
    freq = [0] * 128
    low = 0
    end = 0
    start = 0
    
    while high < len(s):
        windows.add(s[high])
        freq[ord(s[high])] += 1
        
        while len(windows) > k:
            freq[ord(s[low])] -= 1
            if freq[ord(s[low])] == 0:
                windows.remove(s[low])
            
            low += 1
            
        if end - start < high - low:
            end = high
            start = low
        
        high += 1
        
    return s[start:end + 1]

print(getLongest('abcbdbdbbdcdabd',2))
print(getLongest('abcde',2))
print(getLongest('apple',4))
