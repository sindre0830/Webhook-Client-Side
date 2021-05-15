The main information can be found [here](https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2021-workspace/sindre0830/project/weather-events).

### Usage

1. Testing a webhook encryption

  - Input:
    ```
    Method: POST
    Path: .../input/
    ```
    - This is meant for webhook testing
  - Output:
    - Status code and short message
2. Log of webhook tests will show the last 5 inputs where the newest logs are first

  - Input:
    ```
    Method: GET
    Path: .../output/
    ```
  - Output:

    ```go
    type Catch struct {
        Time          string      `json:"time"`
        ErrorMessage  error       `json:"error_message"`
        RawBody       interface{} `json:"raw_body"`
    }
    ```
  - Example:
    - Input:
      ```
      Method: GET
      Path: http://10.212.142.102:8081/client/v1/output/
      ```
    - Output:
      ```json
      Status: 200
      Body:
      [
          {
              "time": "2021-05-15 12:58:28",
              "error_message": null,
              "raw_body": "{\"longitude\":5.33,\"latitude\":60.39,\"location\":\"Bergen, Vestland, Norway\",\"updated\":\"15 May 21 10:50 CEST\",\"date\":\"2021-05-15\",\"data\":{\"instant\":{\"air_temperature\":13.3,\"cloud_area_fraction\":99.1,\"dew_point_temperature\":11.7,\"relative_humidity\":89.4,\"wind_from_direction\":350.4,\"wind_speed\":1.2,\"wind_speed_of_gust\":2.1,\"precipitation_amount\":0},\"predicted\":{\"summary\":\"lightrainshowers_day\",\"confidence\":\"uncertain\",\"air_temperature_max\":16.1,\"air_temperature_min\":12.3,\"precipitation_amount\":0,\"precipitation_amount_max\":1.1,\"precipitation_amount_min\":0,\"probability_of_precipitation\":63.9}}}"
          },
          {
              "time": "2021-05-15 12:58:20",
              "error_message": null,
              "raw_body": "{\"longitude\":2.35,\"latitude\":48.86,\"location\":\"Paris, Ile-de-France, Metropolitan France, France\",\"updated\":\"15 May 21 10:49 CEST\",\"date\":\"2021-05-15\",\"data\":{\"instant\":{\"air_temperature\":9.8,\"cloud_area_fraction\":100,\"dew_point_temperature\":7.7,\"relative_humidity\":87.9,\"wind_from_direction\":170.4,\"wind_speed\":5.9,\"wind_speed_of_gust\":0,\"precipitation_amount\":0},\"predicted\":{\"summary\":\"lightrainshowers_day\",\"confidence\":\"\",\"air_temperature_max\":16.2,\"air_temperature_min\":11.5,\"precipitation_amount\":0.4,\"precipitation_amount_max\":0,\"precipitation_amount_min\":0,\"probability_of_precipitation\":0}}}"
          }
      ]
      ```
