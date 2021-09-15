/*Package api contains base API implementation of unified alerting
 *
 *Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 *
 *Do not manually edit these files, please find ngalert/api/swagger-codegen/ for commands on how to generate them.
 */
package api

import (
	"net/http"

	"github.com/go-macaron/binding"

	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/api/routing"
	"github.com/grafana/grafana/pkg/middleware"
	"github.com/grafana/grafana/pkg/models"
	apimodels "github.com/grafana/grafana/pkg/services/ngalert/api/tooling/definitions"
	"github.com/grafana/grafana/pkg/services/ngalert/metrics"
)

type AlertmanagerApiService interface {
	RouteCreateSilence(*models.ReqContext, apimodels.PostableSilence) response.Response
	RouteDeleteAlertingConfig(*models.ReqContext) response.Response
	RouteDeleteSilence(*models.ReqContext) response.Response
	RouteGetAMAlertGroups(*models.ReqContext) response.Response
	RouteGetAMAlerts(*models.ReqContext) response.Response
	RouteGetAMStatus(*models.ReqContext) response.Response
	RouteGetAlertingConfig(*models.ReqContext) response.Response
	RouteGetSilence(*models.ReqContext) response.Response
	RouteGetSilences(*models.ReqContext) response.Response
	RoutePostAMAlerts(*models.ReqContext, apimodels.PostableAlerts) response.Response
	RoutePostAlertingConfig(*models.ReqContext, apimodels.PostableUserConfig) response.Response
	RoutePostTestReceivers(*models.ReqContext, apimodels.TestReceiversConfigParams) response.Response
}

func (api *API) RegisterAlertmanagerApiEndpoints(srv AlertmanagerApiService, m *metrics.API) {
	api.RouteRegister.Group("", func(group routing.RouteRegister) {
		group.Post(
			toMacaronPath("/api/alertmanager/{Recipient}/api/v2/silences"),
			binding.BindMiddleware(apimodels.PostableSilence{}),
			metrics.Instrument(
				http.MethodPost,
				"/api/alertmanager/{Recipient}/api/v2/silences",
				srv.RouteCreateSilence,
				m,
			),
		)
		group.Delete(
			toMacaronPath("/api/alertmanager/{Recipient}/config/api/v1/alerts"),
			metrics.Instrument(
				http.MethodDelete,
				"/api/alertmanager/{Recipient}/config/api/v1/alerts",
				srv.RouteDeleteAlertingConfig,
				m,
			),
		)
		group.Delete(
			toMacaronPath("/api/alertmanager/{Recipient}/api/v2/silence/{SilenceId}"),
			metrics.Instrument(
				http.MethodDelete,
				"/api/alertmanager/{Recipient}/api/v2/silence/{SilenceId}",
				srv.RouteDeleteSilence,
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/alertmanager/{Recipient}/api/v2/alerts/groups"),
			metrics.Instrument(
				http.MethodGet,
				"/api/alertmanager/{Recipient}/api/v2/alerts/groups",
				srv.RouteGetAMAlertGroups,
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/alertmanager/{Recipient}/api/v2/alerts"),
			metrics.Instrument(
				http.MethodGet,
				"/api/alertmanager/{Recipient}/api/v2/alerts",
				srv.RouteGetAMAlerts,
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/alertmanager/{Recipient}/api/v2/status"),
			metrics.Instrument(
				http.MethodGet,
				"/api/alertmanager/{Recipient}/api/v2/status",
				srv.RouteGetAMStatus,
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/alertmanager/{Recipient}/config/api/v1/alerts"),
			metrics.Instrument(
				http.MethodGet,
				"/api/alertmanager/{Recipient}/config/api/v1/alerts",
				srv.RouteGetAlertingConfig,
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/alertmanager/{Recipient}/api/v2/silence/{SilenceId}"),
			metrics.Instrument(
				http.MethodGet,
				"/api/alertmanager/{Recipient}/api/v2/silence/{SilenceId}",
				srv.RouteGetSilence,
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/alertmanager/{Recipient}/api/v2/silences"),
			metrics.Instrument(
				http.MethodGet,
				"/api/alertmanager/{Recipient}/api/v2/silences",
				srv.RouteGetSilences,
				m,
			),
		)
		group.Post(
			toMacaronPath("/api/alertmanager/{Recipient}/api/v2/alerts"),
			binding.BindMiddleware(apimodels.PostableAlerts{}),
			metrics.Instrument(
				http.MethodPost,
				"/api/alertmanager/{Recipient}/api/v2/alerts",
				srv.RoutePostAMAlerts,
				m,
			),
		)
		group.Post(
			toMacaronPath("/api/alertmanager/{Recipient}/config/api/v1/alerts"),
			binding.BindMiddleware(apimodels.PostableUserConfig{}),
			metrics.Instrument(
				http.MethodPost,
				"/api/alertmanager/{Recipient}/config/api/v1/alerts",
				srv.RoutePostAlertingConfig,
				m,
			),
		)
		group.Post(
			toMacaronPath("/api/alertmanager/{Recipient}/config/api/v1/receivers/test"),
			binding.BindMiddleware(apimodels.TestReceiversConfigParams{}),
			metrics.Instrument(
				http.MethodPost,
				"/api/alertmanager/{Recipient}/config/api/v1/receivers/test",
				srv.RoutePostTestReceivers,
				m,
			),
		)
	}, middleware.ReqSignedIn)
}
