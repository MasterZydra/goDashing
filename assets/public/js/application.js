(function() {
  var lastEvents, source, widgets,
    __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; },
    __bind = function(fn, me){ return function(){ return fn.apply(me, arguments); }; };

  Batman.Filters.prettyNumber = function(num) {
    if (!isNaN(num)) {
      return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    }
  };

  Batman.Filters.dashize = function(str) {
    var dashes_rx1, dashes_rx2;
    dashes_rx1 = /([A-Z]+)([A-Z][a-z])/g;
    dashes_rx2 = /([a-z\d])([A-Z])/g;
    return str.replace(dashes_rx1, '$1_$2').replace(dashes_rx2, '$1_$2').replace(/_/g, '-').toLowerCase();
  };

  Batman.Filters.shortenedNumber = function(num) {
    if (isNaN(num)) {
      return num;
    }
    if (num >= 1000000000) {
      return (num / 1000000000).toFixed(1) + 'B';
    } else if (num >= 1000000) {
      return (num / 1000000).toFixed(1) + 'M';
    } else if (num >= 1000) {
      return (num / 1000).toFixed(1) + 'K';
    } else {
      return num;
    }
  };

  window.Dashing = (function(_super) {

    __extends(Dashing, _super);

    function Dashing() {
      return Dashing.__super__.constructor.apply(this, arguments);
    }

    Dashing.on('reload', function(data) {
      return window.location.reload(true);
    });

    Dashing.root(function() {});

    return Dashing;

  })(Batman.App);

  Dashing.params = Batman.URI.paramsFromQuery(window.location.search.slice(1));

  Dashing.Widget = (function(_super) {

    __extends(Widget, _super);

    function Widget() {
      this.onData = __bind(this.onData, this);

      this.receiveData = __bind(this.receiveData, this);

      var type, _base, _name;
      this.constructor.prototype.source = Batman.Filters.underscore(this.constructor.name);
      Widget.__super__.constructor.apply(this, arguments);
      this.mixin($(this.node).data());
      (_base = Dashing.widgets)[_name = this.id] || (_base[_name] = []);
      Dashing.widgets[this.id].push(this);
      this.mixin(Dashing.lastEvents[this.id]);
      type = Batman.Filters.dashize(this.view);
      $(this.node).addClass("widget widget-" + type + " " + this.id);
    }

    Widget.accessor('updatedAtMessage', function() {
      var hours, minutes, timestamp, updatedAt;
      if (updatedAt = this.get('updatedAt')) {
        timestamp = new Date(updatedAt * 1000);
        hours = timestamp.getHours();
        minutes = ("0" + timestamp.getMinutes()).slice(-2);
        return "Last updated at " + hours + ":" + minutes;
      }
    });

    Widget.prototype.on('ready', function() {
      return Dashing.Widget.fire('ready');
    });

    Widget.prototype.receiveData = function(data) {
      this.mixin(data);
      return this.onData(data);
    };

    Widget.prototype.onData = function(data) {};

    return Widget;

  })(Batman.View);

  Dashing.AnimatedValue = {
    get: Batman.Property.defaultAccessor.get,
    set: function(k, to) {
      var num, num_interval, timer, up,
        _this = this;
      if (!(to != null) || isNaN(to)) {
        return this[k] = to;
      } else {
        timer = "interval_" + k;
        num = !isNaN(this[k]) && (this[k] != null) ? this[k] : 0;
        if (!(this[timer] || num === to)) {
          to = parseFloat(to);
          num = parseFloat(num);
          up = to > num;
          num_interval = Math.abs(num - to) / 90;
          this[timer] = setInterval(function() {
            num = up ? Math.ceil(num + num_interval) : Math.floor(num - num_interval);
            if ((up && num > to) || (!up && num < to)) {
              num = to;
              clearInterval(_this[timer]);
              _this[timer] = null;
              delete _this[timer];
            }
            _this[k] = num;
            return _this.set(k, to);
          }, 10);
        }
        return this[k] = num;
      }
    }
  };

  Dashing.widgets = widgets = {};

  Dashing.lastEvents = lastEvents = {};

  Dashing.debugMode = false;

  source = new EventSource('events');

  source.addEventListener('open', function(e) {
    return console.log("Connection opened", e);
  });

  source.addEventListener('error', function(e) {
    console.log("Connection error", e);
    if (e.currentTarget.readyState === EventSource.CLOSED) {
      console.log("Connection closed");
      return setTimeout((function() {
        return window.location.reload();
      }), 5 * 60 * 1000);
    }
  });

  source.addEventListener('message', function(e) {
    var data, widget, _i, _len, _ref, _ref1, _ref2, _results;
    data = JSON.parse(e.data);
    if (((_ref = lastEvents[data.id]) != null ? _ref.updatedAt : void 0) !== data.updatedAt) {
      if (Dashing.debugMode) {
        console.log("Received data for " + data.id, data);
      }
      lastEvents[data.id] = data;
      if (((_ref1 = widgets[data.id]) != null ? _ref1.length : void 0) > 0) {
        _ref2 = widgets[data.id];
        _results = [];
        for (_i = 0, _len = _ref2.length; _i < _len; _i++) {
          widget = _ref2[_i];
          _results.push(widget.receiveData(data));
        }
        return _results;
      }
    }
  });

  source.addEventListener('dashboards', function(e) {
    var data;
    data = JSON.parse(e.data);
    if (Dashing.debugMode) {
      console.log("Received data for dashboards", data);
    }
    if (data.dashboard === '*' || window.location.pathname === ("/" + data.dashboard)) {
      return Dashing.fire(data.event, data);
    }
  });

  $(document).ready(function() {
    return Dashing.run();
  });

}).call(this);

(function($){
 
    $.fn.extend({ 
         
        leanModal: function(options) {
 
            var defaults = {
                top: 100,
                overlay: 0.5,
                closeButton: null
            }
            
            var overlay = $("<div id='lean_overlay'></div>");
            
            $("body").append(overlay);
                 
            options =  $.extend(defaults, options);
 
            return this.each(function() {
            
                var o = options;
               
                $(this).click(function(e) {
              
                var modal_id = $(this).attr("href");

        $("#lean_overlay").click(function() { 
                     close_modal(modal_id);                    
                });
                
                $(o.closeButton).click(function() { 
                     close_modal(modal_id);                    
                });
                          
                var modal_height = $(modal_id).outerHeight();
              var modal_width = $(modal_id).outerWidth();

            $('#lean_overlay').css({ 'display' : 'block', opacity : 0 });

            $('#lean_overlay').fadeTo(200,o.overlay);

            $(modal_id).css({ 
            
              'display' : 'block',
              'position' : 'fixed',
              'opacity' : 0,
              'z-index': 11000,
              'left' : 50 + '%',
              'margin-left' : -(modal_width/2) + "px",
              'top' : o.top + "px"
            
            });

            $(modal_id).fadeTo(200,1);

                e.preventDefault();
                    
                });
             
            });

      function close_modal(modal_id){

            $("#lean_overlay").fadeOut(200);

            $(modal_id).css({ 'display' : 'none' });
      
      }
    
        }
    });
     
})(jQuery);
/* jshint -W079 */ 

(function() {
  var __bind = function(fn, me){ return function(){ return fn.apply(me, arguments); }; },
    __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  Dashing.Clock = (function(_super) {

    __extends(Clock, _super);

    function Clock() {
      this.startTime = __bind(this.startTime, this);
      return Clock.__super__.constructor.apply(this, arguments);
    }

    Clock.prototype.ready = function() {
      return setInterval(this.startTime, 500);
    };

    Clock.prototype.startTime = function() {
      var h, m, s, today;
      today = new Date();
      h = today.getHours();
      m = today.getMinutes();
      s = today.getSeconds();
      m = this.formatTime(m);
      s = this.formatTime(s);
      this.set('time', h + ":" + m + ":" + s);
      return this.set('date', today.toDateString());
    };

    Clock.prototype.formatTime = function(i) {
      if (i < 10) {
        return "0" + i;
      } else {
        return i;
      }
    };

    return Clock;

  })(Dashing.Widget);

}).call(this);

(function() {
  var __bind = function(fn, me){ return function(){ return fn.apply(me, arguments); }; },
    __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  Dashing.Comments = (function(_super) {

    __extends(Comments, _super);

    function Comments() {
      this.nextComment = __bind(this.nextComment, this);
      return Comments.__super__.constructor.apply(this, arguments);
    }

    Comments.accessor('quote', function() {
      var _ref;
      return "“" + ((_ref = this.get('current_comment')) != null ? _ref.body : void 0) + "”";
    });

    Comments.prototype.ready = function() {
      this.currentIndex = 0;
      this.commentElem = $(this.node).find('.comment-container');
      this.nextComment();
      return this.startCarousel();
    };

    Comments.prototype.onData = function(data) {
      return this.currentIndex = 0;
    };

    Comments.prototype.startCarousel = function() {
      return setInterval(this.nextComment, 8000);
    };

    Comments.prototype.nextComment = function() {
      var comments,
        _this = this;
      comments = this.get('comments');
      if (comments) {
        return this.commentElem.fadeOut(function() {
          _this.currentIndex = (_this.currentIndex + 1) % comments.length;
          _this.set('current_comment', comments[_this.currentIndex]);
          return _this.commentElem.fadeIn();
        });
      }
    };

    return Comments;

  })(Dashing.Widget);

}).call(this);

(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  Dashing.Graph = (function(_super) {

    __extends(Graph, _super);

    function Graph() {
      return Graph.__super__.constructor.apply(this, arguments);
    }

    Graph.accessor('current', function() {
      var points;
      if (this.get('displayedValue')) {
        return this.get('displayedValue');
      }
      points = this.get('points');
      if (points) {
        return points[points.length - 1].y;
      }
    });

    Graph.prototype.ready = function() {
      var container, height, width, x_axis, y_axis;
      container = $(this.node).parent();
      width = (Dashing.widget_base_dimensions[0] * container.data("sizex")) + Dashing.widget_margins[0] * 2 * (container.data("sizex") - 1);
      height = Dashing.widget_base_dimensions[1] * container.data("sizey");
      this.graph = new Rickshaw.Graph({
        element: this.node,
        width: width,
        height: height,
        renderer: this.get("graphtype"),
        series: [
          {
            color: "#fff",
            data: [
              {
                x: 0,
                y: 0
              }
            ]
          }
        ]
      });
      if (this.get('points')) {
        this.graph.series[0].data = this.get('points');
      }
      x_axis = new Rickshaw.Graph.Axis.Time({
        graph: this.graph
      });
      y_axis = new Rickshaw.Graph.Axis.Y({
        graph: this.graph,
        tickFormat: Rickshaw.Fixtures.Number.formatKMBT
      });
      return this.graph.render();
    };

    Graph.prototype.onData = function(data) {
      if (this.graph) {
        this.graph.series[0].data = data.points;
        return this.graph.render();
      }
    };

    return Graph;

  })(Dashing.Widget);

}).call(this);

(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  Dashing.Iframe = (function(_super) {

    __extends(Iframe, _super);

    function Iframe() {
      return Iframe.__super__.constructor.apply(this, arguments);
    }

    Iframe.prototype.ready = function() {};

    Iframe.prototype.onData = function(data) {};

    return Iframe;

  })(Dashing.Widget);

}).call(this);

(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  Dashing.Image = (function(_super) {

    __extends(Image, _super);

    function Image() {
      return Image.__super__.constructor.apply(this, arguments);
    }

    Image.prototype.ready = function() {};

    Image.prototype.onData = function(data) {};

    return Image;

  })(Dashing.Widget);

}).call(this);

(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  Dashing.List = (function(_super) {

    __extends(List, _super);

    function List() {
      return List.__super__.constructor.apply(this, arguments);
    }

    List.prototype.ready = function() {
      if (this.get('unordered')) {
        return $(this.node).find('ol').remove();
      } else {
        return $(this.node).find('ul').remove();
      }
    };

    return List;

  })(Dashing.Widget);

}).call(this);

(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  Dashing.Meter = (function(_super) {

    __extends(Meter, _super);

    Meter.accessor('value', Dashing.AnimatedValue);

    function Meter() {
      Meter.__super__.constructor.apply(this, arguments);
      this.observe('value', function(value) {
        return $(this.node).find(".meter").val(value).trigger('change');
      });
    }

    Meter.prototype.ready = function() {
      var meter;
      meter = $(this.node).find(".meter");
      meter.attr("data-bgcolor", meter.css("background-color"));
      meter.attr("data-fgcolor", meter.css("color"));
      return meter.knob();
    };

    return Meter;

  })(Dashing.Widget);

}).call(this);

(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  Dashing.Number = (function(_super) {

    __extends(Number, _super);

    function Number() {
      return Number.__super__.constructor.apply(this, arguments);
    }

    Number.accessor('current', Dashing.AnimatedValue);

    Number.accessor('difference', function() {
      var current, diff, last;
      if (this.get('last')) {
        last = parseInt(this.get('last'));
        current = parseInt(this.get('current'));
        if (last !== 0) {
          diff = Math.abs(Math.round((current - last) / last * 100));
          return "" + diff + "%";
        }
      } else {
        return "";
      }
    });

    Number.accessor('arrow', function() {
      if (this.get('last')) {
        if (parseInt(this.get('current')) > parseInt(this.get('last'))) {
          return 'icon-arrow-up';
        } else {
          return 'icon-arrow-down';
        }
      }
    });

    Number.prototype.onData = function(data) {
      if (data.status) {
        $(this.get('node')).attr('class', function(i, c) {
          return c.replace(/\bstatus-\S+/g, '');
        });
        return $(this.get('node')).addClass("status-" + data.status);
      }
    };

    return Number;

  })(Dashing.Widget);

}).call(this);

(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  Dashing.Text = (function(_super) {

    __extends(Text, _super);

    function Text() {
      return Text.__super__.constructor.apply(this, arguments);
    }

    return Text;

  })(Dashing.Widget);

}).call(this);

(function() {

  console.log("Yeah! The dashboard has started!");

  Dashing.on('ready', function() {
    var contentWidth;
    Dashing.widget_margins || (Dashing.widget_margins = [5, 5]);
    Dashing.widget_base_dimensions || (Dashing.widget_base_dimensions = [300, 360]);
    Dashing.numColumns || (Dashing.numColumns = 4);
    contentWidth = (Dashing.widget_base_dimensions[0] + Dashing.widget_margins[0] * 2) * Dashing.numColumns;
    return Batman.setImmediate(function() {
      $('.gridster').width(contentWidth);
      return $('.gridster ul:first').gridster({
        widget_margins: Dashing.widget_margins,
        widget_base_dimensions: Dashing.widget_base_dimensions,
        avoid_overlapped_widgets: !Dashing.customGridsterLayout,
        draggable: {
          stop: Dashing.showGridsterInstructions,
          start: function() {
            return Dashing.currentWidgetPositions = Dashing.getWidgetPositions();
          }
        }
      });
    });
  });

}).call(this);
