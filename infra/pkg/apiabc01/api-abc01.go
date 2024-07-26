package apiabc01

type Retorno struct {
	//type DadosGerais struct {
	Tipo  string `json:"tipo"`
	Value string `json:"value"`
}

/* //GetDadosGerais - ...
func GetDadosGerais(stamp, param1, param2 string) (string, error) {
	urlApi := config.AllConfig.APIextrerna + "dadosgerais/" + param1 + "/" + param2
	log.Printf("%v - Vai buscar: %v ", stamp, urlApi)
	response, err := http.Get(urlApi)
	if err != nil {
		log.Printf("%v - ERRO-API: Erro ao buscar: %v ", stamp, urlApi)
		return "", errors.New("Erro ao buscar " + param1 + " -- " + err.Error())
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		//body, err := ioutil.ReadAll(response.Body)
		//if err != nil {
		//	return nil, errors.New("Erro ao buscar " + param1 + " -- " + err.Error())
		//}
		var d *Retorno
		err = json.NewDecoder(response.Body).Decode(&d)
		if err != nil {
			return "", errors.New("Erro ao buscar DadosGerais - " + param1 + " -- " + err.Error())
		}
		return d.Value, nil
	}
	//return "", errors.New("Erro ao buscar " + param1 + " -- StatusCode:" + strconv.Itoa(http.StatusBadRequest))
	return "", errors.New("Erro ao buscar " + param1 + " -- StatusCode:" + strconv.Itoa(response.StatusCode))
} */
