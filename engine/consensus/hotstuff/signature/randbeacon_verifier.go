// +build relic

package signature

import (
	"fmt"

	"github.com/dapperlabs/flow-go/crypto"
	model "github.com/dapperlabs/flow-go/model/hotstuff"
)

// RandomBeaconSigVerifier can verify signatures generated by the random beacon.
// Specifically, it verifies individual key shares and reconstructed threshold signatures.
type RandomBeaconSigVerifier struct {
	randomBeaconHasher     crypto.Hasher     // the hasher for signer random beacon signature
	dkgPubData             *DKGPublicData    // the dkg public data for the only epoch. Should be returned by protocol state if we implement epoch switch
}

func NewRandomBeaconSigVerifier(dkgPubData *DKGPublicData, randomBeaconSigTag string) RandomBeaconSigVerifier {
	return RandomBeaconSigVerifier{
		dkgPubData:             dkgPubData,
		randomBeaconHasher:     crypto.NewBLS_KMAC(randomBeaconSigTag),
	}
}

// VerifyRandomBeaconSig verifies a single random beacon signature for a block using the given public key
// sig - the signature to be verified
// block - the block that the signature was signed for.
// randomBeaconSignerIndex - the signer index of signer's random beacon key share.
func (s *RandomBeaconSigVerifier) VerifyRandomBeaconSig(sig crypto.Signature, block *model.Block, signerPubKey crypto.PublicKey) (bool, error) {
	// convert into message bytes
	msg := BlockToBytesForSign(block)

	// validate random beacon sig with public key
	valid, err := signerPubKey.Verify(sig, msg, s.randomBeaconHasher)
	if err != nil {
		return false, fmt.Errorf("cannot verify random beacon signature: %w", err)
	}

	return valid, nil
}

// VerifyAggregatedRandomBeaconSignature verifies an aggregated random beacon signature, which is a threshold signature
func (s *RandomBeaconSigVerifier) VerifyAggregatedRandomBeaconSignature(sig crypto.Signature, block *model.Block) (bool, error) {
	// convert into bytes
	msg := BlockToBytesForSign(block)

	// the reconstructed signature is also a BLS signature which can be verified by the group public key
	valid, err := s.dkgPubData.GroupPubKey.Verify(sig, msg, s.randomBeaconHasher)
	if err != nil {
		return false, fmt.Errorf("cannot verify reconstructed random beacon sig: %w", err)
	}

	return valid, nil
}
