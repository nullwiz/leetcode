def runLengthEncoding(string):
    ret = ''
    counter = 1
    current_string = string[0] 
    for i in range(1,len(string)):
        if counter == 9 or string[i] != current_string or i == len(string):
            ret += str(counter) + string[i-1]
            counter = 0
            current_string = string[i]
        counter+=1
    return ret + str(counter) + current_string

if __name__ == "__main__":
    print(runLengthEncoding("AAAAAAAAAAAAABBCCCCDD"))
