// const { generateKeyPairSync } = require('crypto');
// const { publicKey, privateKey } = generateKeyPairSync('rsa', {
//   modulusLength: 4096,
//   publicKeyEncoding: {
//     type: 'spki',
//     format: 'pem'
//   },
//   privateKeyEncoding: {
//     type: 'pkcs8',
//     format: 'pem',
//     cipher: 'aes-256-cbc',
//     passphrase: 'top secret'
//   }
// });

var NodeRSA = require('node-rsa');

const key = new NodeRSA({b: 2048});
key.setOptions({ encryptionScheme: 'pkcs1' })
const publicKey = key.exportKey('public')
const privateKey = key.exportKey('private')

export function decrypt(privateKey, encryptBuffer) {
    key.importKey(privateKey)
    const decryptedMsg = key.decrypt(encryptBuffer)
    return decryptedMsg;
}

export function encrypt(publicKey, msg) {
    const key = new NodeRSA(publicKey, 'public')
    const encrypted = key.encrypt(msg)
    return encrypted
}

export {
    key,
    publicKey,
    privateKey
}