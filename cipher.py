def caesar(criptoString, key):
    list1 = list(criptoString)
    for i in range (len(criptoString)):
        if(str(list1[i]) < 'A' or str(list1[i]) > 'Z'):
            list1[i] = criptoString[i]
        else:
            a = int(ord(criptoString[i]))
            list1[i] = chr(((a - ord('A') + key)%26) + ord('A'))
    newString = ''.join(list1)
    print(newString)


def vigenere(criptoString, key):
    z = 0
    list2 = list(criptoString)
    for i in range(len(criptoString)):
        if(z == len(key)):
            z = 0
        if(str(list2[i]) < 'A' or str(list2[i]) > 'Z'):
            list2[i] = criptoString[i]
        else:
            a = int(ord(criptoString[i]))
            list2[i] = chr((a - ord('A') + ord(key[z]) - ord('A'))%26 + ord('A'));
            z = z+1

    newString = ''.join(list2)
    print(newString)


def main():
    op = int(input("Aperte 1 para Cifra de César e 2 para Cifra de Vigenere: "))
    criptoString = input()
    criptoString = criptoString.upper()
    if(op == 1):
        criptoKey1 = int(input())
        caesar(criptoString, criptoKey1)
    if(op == 2):
        criptoKey2 = input()
        criptoKey2 = criptoKey2.upper()
        vigenere(criptoString, criptoKey2)


main()