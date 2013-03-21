class WebTermController extends AppController

  bringToFront: ->
    appStorage = new AppStorage 'WebTerm', '1.0'
    appStorage.fetchStorage =>
      data = new WebTermView appStorage
      data.on "WebTerm.terminated", =>
        @propagateEvent
          KDEventType : "ApplicationWantsToClose"
          globalEvent : yes
        , data: data

      data.on 'ViewClosed', =>
        @propagateEvent
          KDEventType : 'ApplicationWantsToClose'
          globalEvent : yes
        ,
          data : data

      options =
        name         : "Terminal"
        hiddenHandle : no
        type         : "application"
        cssClass     : "webterm"

      @propagateEvent
        KDEventType  : "ApplicationWantsToBeShown"
        globalEvent  : yes
      , {options, data}

WebTerm = {}
