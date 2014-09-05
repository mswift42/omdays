(ns webapp.core
  (:require [om.core :as om :include-macros true]
            [om.dom :as dom :include-macros true]
            [ajax.core :refer [GET POST]]))

(enable-console-print!)

(def app-state (atom {:text "Hello world!"}))

(defn handler [response]
  (.log js/console (str response)))

(GET "/tasks" {:handler handler})



(om/root
  (fn [app owner]
    (reify om/IRender
      (render [_]
        (dom/h1 nil (:text app)))))
  app-state
  {:target (. js/document (getElementById "app"))})
