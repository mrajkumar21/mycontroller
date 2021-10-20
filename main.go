package main

import (
	"flag"
	//"path/filepath"
	"time"

	clientset "mycontroller/pkg/client/clientset/versioned"

	"mycontroller/pkg/signals"

	informers "mycontroller/pkg/client/informers/externalversions"

	//"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	//"k8s.io/client-go/tools/clientcmd"
	//"k8s.io/client-go/util/homedir"
   "k8s.io/client-go/rest"
   "k8s.io/klog/v2"
)

//var kubeconfig *string

func main() {

	stopCh := signals.SetupSignalHandler()

	//if home := homedir.HomeDir(); home != "" {
		//kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	//} else {
		//kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//}
 // creates the connection
 config, err := rest.InClusterConfig()
	if err != nil {
    klog.Errorf("Unable to get cluster config:",err)
		panic(err.Error())
	}
  klog.Infof("Cluster config created")
	flag.Parse()
	// creates the connection
	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	//if err != nil {
		//glog.Fatal(err)
	//}

	// creates the clientset
	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
    klog.Errorf("Unable to create clientset from kubernetes:",err)
		klog.Fatal(err)
	}

	resClient, err := clientset.NewForConfig(config)
	if err != nil {
    klog.Errorf("Unable to create clientset config:",err)
		klog.Fatal(err)
	}
  klog.Info("Clientset creation successful")
	mathInformerFactory := informers.NewSharedInformerFactory(resClient, time.Second*30)
  
	controller := NewController(kubeClient, resClient, mathInformerFactory.Mycontroller().V1alpha1().TestResources())

	// Now let's start the controller

	mathInformerFactory.Start(stopCh)

	if err = controller.Run(2,stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}

}
