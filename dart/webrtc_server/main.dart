import 'dart:convert';
import 'dart:core';
import 'dart:io';

main() async {
  var server = await HttpServer.bind(InternetAddress.loopbackIPv4, 8080);
  print('Serving at ${server.address}:${server.port}');

  await for (var req in server) {
    if (req.uri.path == '/') {
      req.response
        ..headers.contentType = new ContentType('text', 'plain', charset: 'utf-8')
        ..write('Hello, world')
        ..close();

        continue;
    }

    if (req.uri.path == '/signal') {
      var socket = await WebSocketTransformer.upgrade(req);
      socket.listen(
        onDataHandler(socket), 
        onError: onErrorHandler(socket),
        onDone: onDonehandler(socket),
      );

      continue;
    }
  }
}

Map<String, dynamic> users = Map();

bool sendMessage(WebSocket socket, dynamic message) {
    try {
      if (message is String) {
        socket.add(message);
      } else {
        socket.add(jsonEncode(message));
      }

      return true;
    } catch(e) {
      print('[sendMessage] Error $e');
      return false;
    }
  }

dynamic onDataHandler(socket) {
  return (msg) {
    Map<String, dynamic> message = jsonDecode(msg);
    String type = message['type'];

    switch (type) {
      case 'login':
        String name = message['name'];
        print('\t[Message - Login] User($name)');

        if (users.containsKey(name)) {
          print('\t[Message - Login] Remove old User($name)');
          users.remove(name);
        }

        users[name] = socket;
        break;

      default:
        print('[onData] Message($type) from ${message['source']} to ${message['target']}');
        WebSocket target = findUser(message['target']);
        sendMessage(target, message);
    }
  };
}

dynamic onErrorHandler(socket) {
  return (error) {
    var name = null;
    print('[onError]: Connection error ${error.runtimeType}');
    users.forEach((k, v) {
      if (v == socket) {
        name = v;
        print('[onError]: Removed (${k}).');
      }
    });

    if (name != null) {
      users.remove(name);
    }
  };
}

dynamic onDonehandler(socket) {
  return () {
    var name = null;

    users.forEach((k, v) {
      if (v == socket) {
        name = v;
        print('[onDone]: Removed (${k})');
      }
    });

    if (name != null) {
      users.remove(name);
    }
  };
}

WebSocket findUser(String user) {
  return users[user];
}