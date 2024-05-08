package model

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	b64 "encoding/base64"
	"encoding/hex"
)

type Note struct {
    ID int `json:"id" gorm:"primary_key"`
    Notebook string `json:"notebook"`
    Title string `json:"title"`
    Content string `json:"content"`
}

func (n *Note) Decrypt(key string) error {
    hash := md5.Sum([]byte(key))

    aes, err := aes.NewCipher([]byte(hex.EncodeToString(hash[:])))
    if err != nil {
        return err
    }

    gcm, err := cipher.NewGCM(aes)
    if err != nil {
        return err
    }

    title_enc, _ := b64.StdEncoding.DecodeString(n.Title)

    nonceSize := gcm.NonceSize()
    title_nonce, ciphertext := title_enc[:nonceSize], title_enc[nonceSize:]

    title, err := gcm.Open(nil, []byte(title_nonce), []byte(ciphertext), nil)
    if err != nil {
        return err
    }

    n.Title = string(title)

    content_enc, _ := b64.StdEncoding.DecodeString(n.Content)
    
    content_nonce, ciphertext := content_enc[:nonceSize], content_enc[nonceSize:]
    content, err := gcm.Open(nil, []byte(content_nonce), []byte(ciphertext), nil)
    if err != nil {
        return err
    }

    n.Content = string(content)
    return nil
}

func (n *Note) Encrypt(key string) error {
    hash := md5.Sum([]byte(key))

    aes, err := aes.NewCipher([]byte(hex.EncodeToString(hash[:])))
    if err != nil {
        return err
    }

    gcm, err := cipher.NewGCM(aes)
    if err != nil {
        return err
    }

    nonce := make([]byte, gcm.NonceSize())
    _, err = rand.Read(nonce)
    if err != nil {
        return err
    }

    n.Title = b64.StdEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(n.Title), nil))
    n.Content = b64.StdEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(n.Content), nil))
    return nil
}
