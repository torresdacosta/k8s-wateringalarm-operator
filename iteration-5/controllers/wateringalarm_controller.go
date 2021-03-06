/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	alarmv1alpha1 "wateringalarm/api/v1alpha1"
)

// WateringAlarmReconciler reconciles a WateringAlarm object
type WateringAlarmReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=alarm.ricardoptcosta.github.io,resources=wateringalarms,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=alarm.ricardoptcosta.github.io,resources=wateringalarms/status,verbs=get;update;patch

func (r *WateringAlarmReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("wateringalarm", req.NamespacedName)

	log.Info("Hello from Logger", "name", req.NamespacedName)

	var wateringAlarm alarmv1alpha1.WateringAlarm
	if err := r.Get(ctx, req.NamespacedName, &wateringAlarm); err != nil {
		log.Info("error gettinb object", "name", req.NamespacedName)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{}, nil
}

func (r *WateringAlarmReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&alarmv1alpha1.WateringAlarm{}).
		Complete(r)
}
