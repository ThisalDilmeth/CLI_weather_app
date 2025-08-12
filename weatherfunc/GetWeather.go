package weatherfunc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const Grey = "\033[37m"
const Reset = "\033[0m"
const Yellow = "\033[33m"

type weathers struct {
	Name string `json:"name"`

	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`

	Weather []struct {
		Description string `json:"description"`
		ID          int    `json:"id"`
	} `json:"weather"`

	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`

	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`

	Precipitation struct {
		Value float64 `json:"value"`
	} `json:"precipitation"`
}

func Weather(city string) {
	apiKey := os.Getenv("API_KEY")
	url := os.Getenv("API_URL")
	apiUrl := url + city + "&appid=" + apiKey

	res, err := http.Get(apiUrl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather weathers
	err = json.Unmarshal(body, &weather)

	if err != nil {
		panic(err)
	}

	conditionID := weather.Weather[0].ID
	lon, lat := weather.Coord.Lon, weather.Coord.Lat
	temp, feelsLike := weather.Main.Temp, weather.Main.FeelsLike
	pressure, humidity, windSpeed := weather.Main.Pressure, weather.Main.Humidity, weather.Wind.Speed
	location, condition := weather.Name, weather.Weather[0].Description
	precipitationVal := weather.Precipitation.Value
	tempInC := temp - 273.15
	feelsLike = feelsLike - 273.15

	switch {
	case conditionID >= 200 && conditionID <= 232:
		thunderstorm()
	case conditionID >= 300 && conditionID <= 321:
		drizzle()
	case conditionID >= 500 && conditionID <= 531:
		rain()
	case conditionID >= 600 && conditionID <= 622:
		snow()
	case conditionID >= 701 && conditionID <= 781:
		atmosphere()
	case conditionID == 800:
		clear()
	case conditionID >= 801 && conditionID <= 804:
		clouds()
	}

	fmt.Println("\nLocation - " + Yellow + location)
	fmt.Printf(Reset+"Longitude - "+Yellow+"%.3f°\n", lon)
	fmt.Printf(Reset+"Latitude - "+Yellow+"%.3f°\n", lat)
	fmt.Printf(Reset+"Temprature - "+Yellow+"%.2f°C\n", tempInC)
	fmt.Printf(Reset+"FeelsLike - "+Yellow+"%.2f°C\n", feelsLike)
	fmt.Printf(Reset+"Condition - "+Yellow+"%s\n", condition)
	if conditionID >= 500 && conditionID <= 531 || conditionID >= 600 && conditionID <= 622 {
		fmt.Printf(Reset+"Precipitation - "+Yellow+"%.2fmm\n", precipitationVal)
	}
	fmt.Printf(Reset+"Pressure - "+Yellow+"%dhPa\n", pressure)
	fmt.Printf(Reset+"Humidity - "+Yellow+"%d%%\n", humidity)
	fmt.Printf(Reset+"Wind Speed - "+Yellow+"%.2fm/s", windSpeed)
}

func thunderstorm() {
	fmt.Println("\n      .--.     .--.")
	fmt.Println("   .-(    ).-(    ).")
	fmt.Println("  (___.__)___(___.__)")
	fmt.Println("    ⚡ ⚡ ⚡ ⚡ ⚡ ⚡ ⚡ ⚡")
	fmt.Println("   ‘ ‘ ‘ ‘ ‘ ‘ ‘ ‘")
	fmt.Println("  ‘ ‘ ‘ ‘ ‘ ‘ ‘ ‘ ")
}
func drizzle() {
	fmt.Println("\n     .--.")
	fmt.Println("  .-(    ).")
	fmt.Println(" (___.__)__)")
	fmt.Println("     ‘   ‘")
	fmt.Println("    ‘   ‘")
	fmt.Println("   ‘   ‘")
}
func rain() {
	fmt.Println("\n      .--.     .--.")
	fmt.Println("   .-(    ).-(    ).")
	fmt.Println("  (___.__)___(___.__)")
	fmt.Println("     |  |  |  |  |")
	fmt.Println("    |  |  |  |  |  |")
	fmt.Println("   |  |  |  |  |  |  |")
}
func snow() {
	fmt.Println("\n      .--.     .--.")
	fmt.Println("   .-(    ).-(    ).")
	fmt.Println("  (___.__)___(___.__)")
	fmt.Println("     *   *   *   *")
	fmt.Println("   *   *   *   *   *")
	fmt.Println("     *   *   *   *")
}
func atmosphere() {
	fmt.Println("\n           ~   ~   ~    ~    ~   ~   ~")
	fmt.Println("      ~    ~    ~     ~     ~     ~")
	fmt.Println("             ~    ~   _    ~    ~")
	fmt.Println("           ~     .--.( )       ~")
	fmt.Println("      ~        /     (_)          ~")
	fmt.Println("             __/       /    ~   ~     ~")
	fmt.Println("   ~      .-'         '--.      ~")
	fmt.Println("        _(               )_  ~     ~")
	fmt.Println("     .-'   '.__.'`-.__.'   '-.")
	fmt.Println("   (___________________________)")
	fmt.Println("     ~     ~    ~    ~    ~     ~")
	fmt.Println("  ~     ~    ~    ~    ~    ~    ~")
}
func clear() {
	fmt.Println("\n      ;   :   ;")
	fmt.Println("   .   \\_,!,_/   ,")
	fmt.Println("    `.,'     `.,'")
	fmt.Println("     /         \\")
	fmt.Println("~ -- :         : -- ~")
	fmt.Println("     \\         /")
	fmt.Println("    ,'`._   _.'`.")
	fmt.Println("   '   / `!` \\   `")
	fmt.Println("      ;   :   ;")
}
func clouds() {
	fmt.Println("\n       .--.    .--.    .--.")
	fmt.Println("    .-(    ).-(    ).-(    ).")
	fmt.Println("   (___.__)__(___.__)__(___.__)")
	fmt.Println("   .-(    ).-(    ).-(    ).")
	fmt.Println(" (___.__)__(___.__)__(___.__)")
}
