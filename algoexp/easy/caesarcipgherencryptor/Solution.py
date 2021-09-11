def caesarCipherEncryptor(string, key):
    final = ''
    key = key % 26
    for char in string:
        ascinum = ord(char) + key
        if ascinum > 122:
            ascinum = ascinum - 26
        final += chr(ascinum)
    return final


if __name__ == '__main__':
    shifted = caesarCipherEncryptor("abc", 52)
    print(shifted)
