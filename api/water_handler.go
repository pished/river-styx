package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pished/river-styx/model"
	"github.com/pished/river-styx/repository"
	"github.com/pished/river-styx/util/validator"
)

func (a *Api) HandleListWater(w http.ResponseWriter, r *http.Request) {
	conn := DynamoConnect()
	resp, err := repository.ListWater(conn)
	if err != nil {
		a.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrFormDecodingFailure)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (a *Api) HandleAddWater(w http.ResponseWriter, r *http.Request) {
	var form model.WaterForm
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		a.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrFormDecodingFailure)
		return
	}

	if err := a.validator.Struct(form); err != nil {
		a.logger.Warn().Err(err).Msg("")
		resp := validator.ToErrResponse(err)
		if resp == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": "%v"}`, appErrFormErrResponseFailure)
			return
		}
		respBody, err := json.Marshal(resp)
		if err != nil {
			a.logger.Warn().Err(err).Msg("")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": "%v"}`, appErrJsonCreationFailure)
			return
		}

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(respBody)
		return
	}

	waterModel := form.ToModel()
	conn := DynamoConnect()
	err := repository.AddWater(conn, waterModel)
	if err != nil {
		a.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrDataCreationFailure)
	}
	a.logger.Info().Msgf("New item logged: %d", waterModel.Brand)
	w.WriteHeader(http.StatusCreated)
}

func (a *Api) HandleGetWater(w http.ResponseWriter, r *http.Request) {
	conn := DynamoConnect()
	resp, err := repository.GetWater(conn, chi.URLParam(r, "id"))
	if err != nil {
		a.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrFormDecodingFailure)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (a *Api) HandleUpdateWater(w http.ResponseWriter, r *http.Request) {
	var form model.WaterUpdateForm
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		a.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrFormDecodingFailure)
		return
	}

	if err := a.validator.Struct(form); err != nil {
		a.logger.Warn().Err(err).Msg("")
		resp := validator.ToErrResponse(err)
		if resp == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": "%v"}`, appErrFormErrResponseFailure)
			return
		}
		respBody, err := json.Marshal(resp)
		if err != nil {
			a.logger.Warn().Err(err).Msg("")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": "%v"}`, appErrJsonCreationFailure)
			return
		}

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(respBody)
		return
	}

	waterModel := form.ToModel()
	conn := DynamoConnect()
	err := repository.UpdateWater(conn, waterModel, chi.URLParam(r, "id"))
	if err != nil {
		a.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrDataCreationFailure)
	}
	a.logger.Info().Msgf("Upated: %d", waterModel.Brand)
	w.WriteHeader(http.StatusAccepted)
}

func (a *Api) HandleDeleteWater(w http.ResponseWriter, r *http.Request) {
	conn := DynamoConnect()
	err := repository.DeleteWater(conn, chi.URLParam(r, "id"))
	if err != nil {
		a.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrFormDecodingFailure)
		return
	}
	a.logger.Info().Msgf("Deleted: %d", chi.URLParam(r, "id"))
	w.WriteHeader(http.StatusAccepted)
}
