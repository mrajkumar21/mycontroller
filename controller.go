package main

import (
	"fmt"
 "context"
 metav1"k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
  "reflect"
	//"github.com/golang/glog"
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
 "k8s.io/klog/v2"
  samplecontroller "mycontroller/pkg/apis/samplecontroller/v1alpha1"
	clientset "mycontroller/pkg/client/clientset/versioned"
	mathresourcescheme "mycontroller/pkg/client/clientset/versioned/scheme"
	informers "mycontroller/pkg/client/informers/externalversions/samplecontroller/v1alpha1"
	listers "mycontroller/pkg/client/listers/samplecontroller/v1alpha1"
)
var entry int
const(
   SuccessStatus = "Success"
   SuccessMessage = "Updated Successfully"
)
const controllerAgentName = "mycontroller"

type Controller struct {
	kubeclientset kubernetes.Interface

	resclientset clientset.Interface

	testresourcesLister listers.TestResourceLister
	testresourcesSynced cache.InformerSynced

	workqueue workqueue.RateLimitingInterface
	informer  cache.SharedIndexInformer

	recorder record.EventRecorder
}

func NewController(
	kubeclientset kubernetes.Interface, resclientset clientset.Interface,
	testResourceInformer informers.TestResourceInformer) *Controller {
  klog.Info("<<<<<In Controller>>>>>")
	utilruntime.Must(mathresourcescheme.AddToScheme(scheme.Scheme))
	klog.Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(klog.Infof)
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

	klog.Info("Setting up event handlers")
	// Set up an event handler for when Student resources change
	testResourceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueTestResource,
		UpdateFunc: func(old, new interface{}) {
      newMath := new.(*samplecontroller.TestResource)
        oldMath := old.(*samplecontroller.TestResource)
        if reflect.DeepEqual(newMath.Spec, oldMath.Spec) {
        klog.Info("Status:",newMath.Status)
        klog.Info("Specs not modified. Ignoring update event")
        return
       }
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
    err:= func(obj interface{})error {
		defer c.workqueue.Done(obj)
		var key string
		var ok bool

		if key, ok = obj.(string); !ok {

			c.workqueue.Forget(obj)
			utilruntime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}
		if err := c.syncHandler(key); err != nil {
      klog.Errorf("Error in SyncHandler:",err)
			c.workqueue.AddRateLimited(key)
			return fmt.Errorf("error syncing '%s': %s, requeing", key, err.Error())
		}
		c.workqueue.Forget(obj)
		klog.Infof("Successfully synced '%s'", key)
		return nil
  }(obj)
	if err != nil {
      klog.Errorf("Error in :",err)
      utilruntime.HandleError(err)
		  return true
	}

	return true
}

func (c *Controller) syncHandler(key string) error {

	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
    klog.Errorf("Invalid resource key:",err)
		return nil
	}

	cmath, err := c.testresourcesLister.TestResources(namespace).Get(name)
	if err != nil {
		klog.Errorf("Fetching CRD  with key %s from store failed with %v", key, err)
		return err
	}

	if cmath.Spec.Operation != "" {

		switch cmath.Spec.Operation {

		case ("add"):
			{
				klog.Infof("Operation Addition  value= %d \n", cmath.Spec.FirstNum+ cmath.Spec.SecondNum)

			}

		case ("sub"):
			{
				klog.Infof("Operation subtraction value= %d \n", cmath.Spec.FirstNum- cmath.Spec.SecondNum)

			}
		case ("mul"):
			{
				klog.Infof("Operation multiplication  value= %d \n", cmath.Spec.FirstNum* cmath.Spec.SecondNum)

			}

		case ("div"):
			{
				klog.Infof("Operation division value= %d \n", cmath.Spec.FirstNum/ cmath.Spec.SecondNum)

			}

		}

	} else {

		klog.Errorf("Fetching object cmath.Spec.Operation with  key %s from store failed with %v", key, err)
		return err

	}
  err = c.updateSamplecontrollerStatus(cmath)
   if err != nil {
	 klog.Fatal(err)
  }
	return nil

}
func (c *Controller) updateSamplecontrollerStatus(foo *samplecontroller.TestResource) error {

  fooCopy := foo.DeepCopy()
  fooCopy.Status.State = SuccessStatus
  fooCopy.Status.Message = SuccessMessage
  fooCopy.Status.Created_At= time.Now()
  _, err := c.resclientset.MycontrollerV1alpha1().TestResources(foo.Namespace).UpdateStatus(context.TODO(), fooCopy, metav1.UpdateOptions{})
  
  return err
}
func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()

	// Let the workers stop when we are done
	defer c.workqueue.ShutDown()
	klog.Info("start controller Business, start a cache data synchronization")
	if ok := cache.WaitForCacheSync(stopCh, c.testresourcesSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	klog.Info("worker start-up")
	for i := 0; i < threadiness; i++ {
		 go wait.Until(c.runWorker, time.Second, stopCh)
	}
        
        klog.Info("worker Already Started")
        <-stopCh
      	klog.Info("worker It's already over.")

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
    klog.Errorf("Error in cache objects:",err)
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
    klog.Errorf("Delete Operation Failed:",err)
		runtime.HandleError(err)
		return
	}
	//Queue the key again
	c.workqueue.AddRateLimited(key)
}
