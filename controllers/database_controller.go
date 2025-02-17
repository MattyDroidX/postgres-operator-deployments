package controllers

import (
	"context"

	v1 "github.com/example/postgres-operator/api/v1"
	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// DatabaseReconciler reconciles a Database object
type DatabaseReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=example.com,resources=databases,verbs=get;list;watch;create;update;patch;delete

func (r *DatabaseReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("database", req.NamespacedName)

	// Fetch the Database instance
	var database v1.Database
	if err := r.Get(ctx, req.NamespacedName, &database); err != nil {
		log.Error(err, "Unable to fetch Database")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Create a PostgreSQL deployment if it doesn’t exist
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      database.Name,
			Namespace: database.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: ptrInt32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": database.Name},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"app": database.Name}},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "postgres",
							Image: "postgres:" + database.Spec.Version,
							Env: []corev1.EnvVar{
								{Name: "POSTGRES_PASSWORD", Value: "example"},
							},
						},
					},
				},
			},
		},
	}

	if err := r.Create(ctx, deployment); err != nil {
		log.Error(err, "Failed to create Deployment")
		return ctrl.Result{}, err
	}

	database.Status.Status = "Provisioned"
	if err := r.Status().Update(ctx, &database); err != nil {
		log.Error(err, "Failed to update Database status")
		return ctrl.Result{}, err
	}

	log.Info("Database provisioned successfully", "Database Name", database.Name)
	return ctrl.Result{}, nil
}

func ptrInt32(i int32) *int32 { return &i }

func (r *DatabaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.Database{}).
		Complete(r)
}
