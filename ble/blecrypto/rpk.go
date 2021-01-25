package blecrypto

import (
	"github.com/digital-dream-labs/vector-bluetooth/rts"
	"github.com/jamesruan/sodium"
)

// SetRemotePublicKey populates the remote public key
func (b *BLECrypto) SetRemotePublicKey(msg *rts.RtsConnRequest) error {
	b.remotePublicKey = sodium.KXPublicKey{
		Bytes: msg.PublicKey[:],
	}
	return nil
}

// GetRemotePublicKey returns the public key
func (b *BLECrypto) GetRemotePublicKey() [32]uint8 {
	pkb := [32]uint8{}
	for k, v := range b.keys.PublicKey.Bytes {
		pkb[k] = v
	}
	return pkb
}
