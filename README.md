kind: glotPod

group: goglot
version: v1alpha1

Three fields are required for a k8 object. 
Objectmeta
TypeMeta
Spec


For spec I would need the things like-
Language-
Input-
Code-
Name-

Then we would need to register our types in k8s schemes

For out crd object, we would need to have a DeepCopyObject, we would need to have our own clienset, informers and listers. We will use code-generator for it.
