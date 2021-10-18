package main

import (
	"fmt"
	"time"

	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"

	clientset "github.com/mrajkumar21/mycontroller/pkg/client/clientset/versioned"
	mathresourcescheme "github.com/mrajkumar21/mycontroller/pkg/client/clientset/versioned/scheme"
	informers "github.com/mrajkumar21/mycontroller/pkg/client/informers/externalversions/Mycontroller/v1alpha1"
	listers "github.com/mrajkumar21/mycontroller/pkg/client/listers/mycontroller/v1alpha1"
)

const controllerAgentName = "mycontroller"

type Controller struct {
	kubeclientset kubernetes.Interface

	resclientset clientset.Interface

	testresourcesLister listers.MathResourceLister
	testresourcesSynced cache.InformerSynced

	workqueue workqueue.RateLimitingInterface
	informer  cache.SharedIndexInformer

	recorder record.EventRecorder
}

func NewController(
	kubeclientset kubernetes.Interface, resclientset clientset.Interface,
	testResourceInformer informers.MathResourceInformer) *Controller {

	utilruntime.Must(mathresourcescheme.AddToScheme(scheme.Scheme))
	glog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	controller := &Controller{
		kubeclientset:       kubeclientset,
		resclientset:        resclientset,
		testresourcesLister: testResourceInformer.Lister(),
		testresourcesSynced: testResourceInformer.Informer().HasSynced,
		workqueue:           workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "TestResource"),
		recorder:            recorder,
	}

	glog.Info("Setting up event handlers")
	// Set up an event handler for when Student resources change
	testResourceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueTestResource,
		UpdateFunc: func(old, new interface{}) {
			controller.enqueueTestResource(new)
		},
		DeleteFunc: controller.enqueueTestResourceForDelete,
	})

	return controller
}

func (c *Controller) processNextItem() bool {
	obj, shutdown := c.workqueue.Get()

	if shutdown {
		return false
	}

	err := func(obj interface{}) error {
		defer c.workqueue.Done(obj)
		var key string
		var ok bool

		if key, ok = obj.(string); !ok {

			c.workqueue.Forget(obj)
			utilruntime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}
		if err := c.syncHandler(key); err != nil {
			c.workqueue.AddRateLimited(key)
			return fmt.Errorf("error syncing '%s': %s, requeuing", key, err.Error())
		}
		c.workqueue.Forget(obj)
		glog.Infof("Successfully synced '%s'", key)
		return nil
	}(obj)

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}

	return true
}

func (c *Controller) syncHandler(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	cmath, err := c.testresourcesLister.MathResources(namespace).Get(name)
	if err != nil {
		glog.Errorf("Fetching CRD  with key %s from store failed with %v", key, err)
		return err
	}

	if cmath.Spec.Operation != "" {

		switch cmath.Spec.Operation {

		case ("add"):
			{
				fmt.Printf("Operation Addition  value= %d \n", cmath.Spec.FirstNum+cmath.Spec.SecondNum)

			}

		case ("sub"):
			{
				fmt.Printf("Operation subtraction value= %d \n", cmath.Spec.FirstNum-cmath.Spec.SecondNum)

			}
		case ("mul"):
			{
				fmt.Printf("Operation multiplication  value= %d \n", cmath.Spec.FirstNum*cmath.Spec.SecondNum)

			}

		case ("div"):
			{
				fmt.Printf("Operation division value= %d \n", cmath.Spec.FirstNum/cmath.Spec.SecondNum)

			}

		}

	} else {

		glog.Errorf("Fetching object cmath.Spec.Operation with  key %s from store failed with %v", key, err)
		return err

	}

	return nil

}

func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()

	// Let the workers stop when we are done
	defer c.workqueue.ShutDown()
	glog.Info("start controller Business, start a cache data synchronization")
	if ok := cache.WaitForCacheSync(stopCh, c.testresourcesSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	glog.Info("worker start-up")
	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	glog.Info("worker Already started")
	<-stopCh
	glog.Info("worker It's already over.")

	return nil

}

func (c *Controller) runWorker() {
	for c.processNextItem() {
	}
}

func (c *Controller) enqueueTestResource(obj interface{}) {
	var key string
	var err error
	// Cache objects
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		runtime.HandleError(err)
		return
	}

	// Queue key s
	c.workqueue.AddRateLimited(key)
}

// Delete operation
func (c *Controller) enqueueTestResourceForDelete(obj interface{}) {
	var key string
	var err error
	// Delete the specified object from the cache
	key, err = cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(err)
		return
	}
	//Queue the key again
	c.workqueue.AddRateLimited(key)
}
