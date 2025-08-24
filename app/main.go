package main

import (
    "flag"
    "log"
    "time"

    v1 "k8s.io/api/core/v1"
    "k8s.io/client-go/informers"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
    "k8s.io/client-go/tools/cache"
)

func main() {
    verbose := flag.Bool("v", false, "Enable verbose logging")
    flag.Parse()

    // Cluster içinden bağlan
    config, err := rest.InClusterConfig()
    if err != nil {
        log.Fatalf("Error creating in-cluster config: %v", err)
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        log.Fatalf("Error creating clientset: %v", err)
    }

    factory := informers.NewSharedInformerFactoryWithOptions(
        clientset,
        time.Minute,
        informers.WithNamespace("jenkins"),
    )

    podInformer := factory.Core().V1().Pods().Informer()

    podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
        AddFunc: func(obj interface{}) {
            pod := obj.(*v1.Pod)
            if *verbose {
                log.Printf("New Pod detected: %s/%s, phase: %s", pod.Namespace, pod.Name, pod.Status.Phase)
            }

            if pod.Labels["job-name"] != "" {
                log.Printf("Jenkins job pod started: %s/%s (job: %s)", pod.Namespace, pod.Name, pod.Labels["job-name"])
            }
        },
        UpdateFunc: func(oldObj, newObj interface{}) {
            oldPod := oldObj.(*v1.Pod)
            newPod := newObj.(*v1.Pod)

            if oldPod.Status.Phase != newPod.Status.Phase {
                log.Printf("Pod updated: %s/%s, phase: %s -> %s", newPod.Namespace, newPod.Name, oldPod.Status.Phase, newPod.Status.Phase)
            }
        },
        DeleteFunc: func(obj interface{}) {
            pod := obj.(*v1.Pod)
            log.Printf("Pod deleted: %s/%s", pod.Namespace, pod.Name)
        },
    })

    stopCh := make(chan struct{})
    defer close(stopCh)

    factory.Start(stopCh)

    if !cache.WaitForCacheSync(stopCh, podInformer.HasSynced) {
        log.Fatalf("Failed to sync caches")
    }

    log.Println("✅ Test Pod informer started. Watching Jenkins pods...")
    <-stopCh
}
