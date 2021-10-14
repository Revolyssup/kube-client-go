package main

/*****For Operator*******/
import (
	"fmt"

	op "github.com/Revolyssup/kube-client-go/pkg/api/goglot.dev/v1alpha1"
)

func main() {
	k := op.GlotPod{}
	fmt.Println(k)
}

/*******For goglot server, this might be helpful**********/
// func main() {
// 	var config *rest.Config
// 	config, err := clientcmd.BuildConfigFromFlags("", "/home/ashish/.kube/config")
// 	if err != nil {
// 		config, err = rest.InClusterConfig()
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	fmt.Println("bt ", config.BearerToken)
// 	clientset := kubernetes.NewForConfigOrDie(config)
// 	var stop = make(chan struct{})
// 	go Watcher(clientset, stop)
// 	newPod := &corev1.Pod{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name: "test-pod",
// 			Labels: map[string]string{
// 				"app": "revoly",
// 			},
// 		},
// 		Spec: corev1.PodSpec{
// 			Containers: []corev1.Container{{Name: "cont", Image: "revoly/portfolio", Ports: []corev1.ContainerPort{{HostPort: 3000, ContainerPort: 80}}}},
// 		},
// 	}
// 	newSvc := &corev1.Service{
// 		TypeMeta: metav1.TypeMeta{},
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name: "test-svc",
// 		},
// 		Spec: corev1.ServiceSpec{
// 			Selector: map[string]string{
// 				"app": "revoly",
// 			},
// 			Ports: []corev1.ServicePort{{
// 				Protocol:   corev1.ProtocolTCP,
// 				Port:       3000,
// 				TargetPort: intstr.FromInt(3000),
// 			}},
// 			Type: corev1.ServiceTypeLoadBalancer,
// 		},
// 	}
// 	pod, err := clientset.CoreV1().Pods("default").Create(context.TODO(), newPod, metav1.CreateOptions{})

// 	if err != nil {
// 		fmt.Println("[Err] ", err.Error())
// 	}
// 	fmt.Println(pod.Name + " pod created!!")

// 	svc, err := clientset.CoreV1().Services("default").Create(context.Background(), newSvc, metav1.CreateOptions{})
// 	if err != nil {
// 		fmt.Println("[Err] ", err.Error())
// 	}
// 	fmt.Println(svc.Name + " service created!!")
// 	select {
// 	case <-stop:
// 		fmt.Println("STOPPED watching")
// 	}
// }
// func Watcher(clientset *kubernetes.Clientset, stop chan struct{}) {
// 	fmt.Println("HERE")
// 	inf := informers.NewFilteredSharedInformerFactory(clientset, time.Minute, "default", nil)
// 	podInformer := inf.Core().V1().Pods()
// 	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
// 		AddFunc: func(obj interface{}) {
// 			labels := obj.(*v1.Pod).GetLabels()
// 			fmt.Println("app set to ", labels["app"])
// 		},
// 	})
// 	go podInformer.Informer().Run(stop)
// 	svcInformer := inf.Core().V1().Endpoints()
// 	svcInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
// 		AddFunc: func(obj interface{}) {
// 			fmt.Println("I saw an endpoint ", obj.(*v1.Endpoints).GetName(), "\nAlso name is : ", obj.(*v1.Endpoints).Name)
// 		},
// 	})
// 	go svcInformer.Informer().Run(stop)
// }
