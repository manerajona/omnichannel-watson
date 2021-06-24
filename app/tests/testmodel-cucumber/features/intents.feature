Feature: Infoguia Intents
  In order to determine if the assistant is behaving correctly
  We need to check the following utterances

  Scenario:
    When user asks to watson: "menu"
    Then watson should respond with intent: "ver_menu"

  Scenario:
    When user asks to watson: "comercios"
    Then watson should respond with intent: "ver_comercios"
  
  Scenario:
    When user asks to watson: "animales"
    Then watson should respond with intent: "ver_animales"

  Scenario:
    When user asks to watson: "automovil"
    Then watson should respond with intent: "ver_automovil"

  Scenario:
    When user asks to watson: "construccion"
    Then watson should respond with intent: "ver_construccion"

  Scenario:
    When user asks to watson: "deportes"
    Then watson should respond with intent: "ver_deportes"

  Scenario:
    When user asks to watson: "educacion"
    Then watson should respond with intent: "ver_educacion"

  Scenario:
    When user asks to watson: "entretenimiento"
    Then watson should respond with intent: "ver_entretenimiento"

  Scenario:
    When user asks to watson: "gastronomia"
    Then watson should respond with intent: "ver_gastronomia"

  Scenario:
    When user asks to watson: "hogar"
    Then watson should respond with intent: "ver_hogar"

  Scenario:
    When user asks to watson: "indumentaria"
    Then watson should respond with intent: "ver_indumentaria"

  Scenario:
    When user asks to watson: "profesionales"
    Then watson should respond with intent: "ver_profesionales"

  Scenario:
    When user asks to watson: "regaleria"
    Then watson should respond with intent: "ver_regaleria"

  Scenario:
    When user asks to watson: "estetica"
    Then watson should respond with intent: "ver_salud_estetica"

  Scenario:
    When user asks to watson: "viajes"
    Then watson should respond with intent: "ver_viajes"

  Scenario:
    When user asks to watson: "emergencias"
    Then watson should respond with intent: "ver_emergencias"

  Scenario:
    When user asks to watson: "entidades"
    Then watson should respond with intent: "ver_entidades"

  Scenario:
    When user asks to watson: "bancos"
    Then watson should respond with intent: "ver_bancos"

  Scenario:
    When user asks to watson: "estatales"
    Then watson should respond with intent: "ver_estatales"

  Scenario:
    When user asks to watson: "escuelas"
    Then watson should respond with intent: "ver_escuelas"

  Scenario:
    When user asks to watson: "hospitales"
    Then watson should respond with intent: "ver_hospitales"

  Scenario:
    When user asks to watson: "servicios esenciales"
    Then watson should respond with intent: "ver_servicios_esenciales"

  Scenario:
    When user asks to watson: "farmacias"
    Then watson should respond with intent: "ver_farmacias"

  Scenario:
    When user asks to watson: "remis"
    Then watson should respond with intent: "ver_remis"

  Scenario:
    When user asks to watson: "eventos"
    Then watson should respond with intent: "ver_eventos"

  Scenario:
    When user asks to watson: "ofertas"
    Then watson should respond with intent: "ver_ofertas"

  Scenario:
    When user asks to watson: "turismo"
    Then watson should respond with intent: "ver_lugares_turisticos"

  Scenario:
    When user asks to watson: "actividades"
    Then watson should respond with intent: "ver_actividades"

  Scenario:
    When user asks to watson: "bibliotecas"
    Then watson should respond with intent: "ver_bibliotecas"

  Scenario:
    When user asks to watson: "cementerios"
    Then watson should respond with intent: "ver_cementerios"

  Scenario:
    When user asks to watson: "museos"
    Then watson should respond with intent: "ver_museos"

  Scenario:
    When user asks to watson: "parques"
    Then watson should respond with intent: "ver_plazas_parques"

  Scenario:
    When user asks to watson: "templos"
    Then watson should respond with intent: "ver_templos"

  Scenario:
    When user asks to watson: "terminales"
    Then watson should respond with intent: "ver_terminales"