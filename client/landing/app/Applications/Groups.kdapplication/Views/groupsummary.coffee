class GroupSummaryView extends KDCustomHTMLView

  constructor:(options = {}, data = {})->

    options.domId = "group-summary"

    super options, data

    @lazyDomController = @getSingleton("lazyDomController")

    @lazyDomController.on "landingViewIsShown",  @bound "landingViewIsShown"
    @lazyDomController.on "landingViewIsHidden", @bound "landingViewIsHidden"

    @loader = new KDLoaderView
      size          :
        width       : 60
      loaderOptions :
        color       : "#ffffff"


    @sign = new KDCustomHTMLView
      tagName     : "div"
      cssClass    : "group-sign-wrapper"
      bind        : "mouseenter mouseleave"
      pistachio   : """
        <i></i><i></i>
        <h3 id="group-sign">{{#(title)}}</h3>
        """
      mouseenter  : -> @$().css marginTop : 13
      mouseleave  : -> @$().css marginTop : 8
      click       : @bound "showSummary"
    , {}

    @sign.bindTransitionEnd()

    @kodingLogo = new KDCustomHTMLView
      tagName     : "div"
      cssClass    : "summary-koding-logo"
      click       : (event)=>
        @kodingLogo.unsetClass "in"
        @kodingLogo.once "transitionend", =>
          @showSummary event


    @kodingLogo.bindTransitionEnd()

    @utils.wait 1500, => @kodingLogo.setClass "in"

    @landingPageLink = new CustomLinkView
      cssClass    : "public-page-link"
      title       : "Pull up the public page"
      icon        :
        placement : "right"
      click       : @bound "showLandingPage"

    @openInKodingLink = new CustomLinkView
      cssClass    : "open-in-koding"
      title       : "Open the group in Koding"
      icon        : {}
      click       : @bound "hideLandingPage"

    @closeLink = new CustomLinkView
      cssClass    : "close-link"
      title       : ""
      icon        :
        placement : "right"
      click       : (event)=>
        event.stopPropagation()
        event.preventDefault()
        @hideSummary()
        @utils.wait 400, =>
          @sign.setClass "swing-in"

    @once "viewAppended", @loader.bound "show"
    @once "viewAppended", @bound "decorateSummary"
    @bindTransitionEnd()

  viewAppended: JView::viewAppended

  pistachio:->
    """
      {{> @kodingLogo}}
      {{> @loader}}
      <header>
        <div class='avatar-wrapper'></div>
        <div class='right-overflow'></div>
      </header>
      <aside></aside>
      {{> @landingPageLink}}
      {{> @closeLink}}
      {{> @openInKodingLink}}
      {{> @sign}}
    """

  showSummary:(event)->
    if event
      event.stopPropagation()
      event.preventDefault()
    if @lazyDomController.isLandingPageVisible()
      @$().css top : -@getHeight()
    else
      @sign.once "transitionend", =>
        @sign.unsetClass "swing-out"

      @$().css top : 0
      @setClass "swing-out"
      @unsetClass "swing-in"


  hideSummary:(event)->
    if event
      event.stopPropagation()
      event.preventDefault()
    if @lazyDomController.isLandingPageVisible()
      @$().css top : 0
      @once "transitionend", =>
        @kodingLogo.setClass "in"
    else
      @once "transitionend", =>
        @sign.unsetClass "swing-out"
        @sign.unsetClass "swing-in"
      @$().css top : -@getHeight()

  landingViewIsHidden:->

    @sign.setClass "swing-in"

  landingViewIsShown:->

    @sign.unsetClass "swing-in swing-out"

  decorateSummary:->

    KD.remote.cacheable KD.config.groupEntryPoint, (err, models)=>
      if err then callback err
      else if models?
        [group] = models
        @sign.setData group
        @sign.render()

        group.fetchAdmin (err, owner)=>
          if err then warn err
          else
            @loader.hide()
            @putOwner owner
            @putGroupBio group

        group.fetchMembers (err, members)=>
          if err then warn err
          else if members.length
            @putList members

  putOwner:(owner)->

    ownerAvatar = new AvatarView
      cssClass  : "owner"
      size      :
        width   : 60
        height  : 60
    , owner

    @addSubView ownerAvatar, '.avatar-wrapper'

    title = new KDHeaderView
      type  : "medium"
      title : "#{@utils.getFullnameFromAccount owner} created this group."

    @addSubView title, '.right-overflow'

  putGroupBio:(group)->

    bio = new KDCustomHTMLView
      tagName   : "p"
      cssClass  : "bio"
      partial   : group.body or "This is a group created within Koding, group admins decide wether to share members/activities/topics of this group. Unless it is wanted Koding members won't be able to see the content shared in this group."

    @addSubView bio, '.right-overflow'

  putList:(members)->

    controller = new KDListViewController
      view         : @members = new KDListView
        wrapper    : no
        scrollView : no
        type       : "members"
        itemClass  : GroupItemMemberView

    @addSubView @members, 'aside'
    controller.instantiateListItems members

  showLandingPage:(event)->

    event.stopPropagation()
    event.preventDefault()
    @lazyDomController.showLandingPage =>
      @hideSummary()

  hideLandingPage:(event)->

    event.stopPropagation()
    event.preventDefault()
    @lazyDomController.hideLandingPage()