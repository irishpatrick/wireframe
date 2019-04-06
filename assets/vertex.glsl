#version 330 core

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec3 vert;
in vec2 vertTexCoord;
in vec3 norm;

out vec3 frag;
out vec3 normal;
out vec2 fragTexCoord;

void main() {
    fragTexCoord = vertTexCoord;
    gl_Position = projection * camera * model * vec4(vert, 1);
    normal = vec3(model * vec4(norm, 1));
    frag = vec3(model * vec4(vert, 1.0));
}
