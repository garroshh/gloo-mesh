package v1

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
)

var _ = Describe("TranslatorEventLoop", func() {
	var (
		namespace string
		emitter     TranslatorEmitter
		err       error
	)

	BeforeEach(func() {

		meshClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		meshClient, err := NewMeshClient(meshClientFactory)
		Expect(err).NotTo(HaveOccurred())

		emitter = NewTranslatorEmitter(meshClient)
	})
	It("runs sync function on a new snapshot", func() {
		_, err = emitter.Mesh().Write(NewMesh(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		sync := &mockTranslatorSyncer{}
		el := NewTranslatorEventLoop(emitter, sync)
		_, err := el.Run([]string{namespace}, clients.WatchOpts{})
		Expect(err).NotTo(HaveOccurred())
		Eventually(func() bool { return sync.synced }, time.Second).Should(BeTrue())
	})
})

type mockTranslatorSyncer struct {
	synced bool
}

func (s *mockTranslatorSyncer) Sync(ctx context.Context, snap *TranslatorSnapshot) error {
	s.synced = true
	return nil
}
