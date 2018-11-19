module Main exposing (main)

import Browser
import Html
import Html.Events as Events


main : Program Bool Bool Msg
main =
    Browser.application
        { init = init
        , view = view
        , update = update
        , subscriptions = always Sub.none
        , onUrlChange = always NoOp
        , onUrlRequest = always NoOp
        }


init flags _ _ =
    ( flags, Cmd.none )


type Msg
    = NoOp
    | Toggle


update msg state =
    case msg of
        NoOp ->
            ( state, Cmd.none )

        Toggle ->
            ( not state, Cmd.none )


view model =
    { title = ""
    , body =
        [ Html.div []
            [ Html.text <|
                if model then
                    "On"

                else
                    "Off"
            ]
        , Html.button
            [ Events.onClick Toggle ]
            [ Html.text "Toggle" ]
        ]
    }
