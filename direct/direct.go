package direct

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/skycoin/dmsg"
	"github.com/skycoin/dmsg/cipher"
)

// StartDmsg
func StartDmsg(ctx context.Context, log logrus.FieldLogger, pk cipher.PubKey, sk cipher.SecKey) (dmsgC *dmsg.Client, stop func(), err error) {
	dmsgC = dmsg.NewClient(pk, sk, NewClient(GetAllEntries(pk)), dmsg.DefaultConfig())
	go dmsgC.Serve(context.Background())

	stop = func() {
		err := dmsgC.Close()
		log.WithError(err).Info("Disconnected from dmsg network.")
	}

	log.WithField("public_key", pk.String()).
		Info("Connecting to dmsg network...")

	select {
	case <-ctx.Done():
		stop()
		return nil, nil, ctx.Err()

	case <-dmsgC.Ready():
		log.Info("Dmsg network ready.")
		return dmsgC, stop, nil
	}
}
