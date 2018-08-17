# -*- encoding: utf-8 -*-
# that script aim to convert VCARD format to CSV

import sys

def main(filename):


    with open(filename) as vcf :
        for line in vcf:
            #print(line)
            line = line[:-1]
            if line == "BEGIN:VCARD":
                sys.stdout.write("\n")
            fields = line.split(":")
            #print(fields)
            if fields[0] == "N" :

                sys.stdout.write(","+fields[1])
            if fields[0] == "TEL;HOME":
                sys.stdout.write(",\""+fields[1]+"\"")



if __name__ == "__main__":
    main(sys.argv[1])
