#version 330 core

uniform sampler2D tex;

in vec3 frag;
in vec3 normal;
in vec2 fragTexCoord;

out vec4 outputColor;

void main() {
    vec3 light = vec3(1, 1, 0);
    float intensity = 0.8;
    light = light * intensity;
    vec3 norm = normalize(normal);
    vec3 dir = normalize(light - frag);

    float diff = max(dot(norm, light), 0.0);
    vec3 diffuse = diff * vec3(1,1,1);

    float ambient = 0.0;
    vec4 tex = texture(tex, fragTexCoord);
    vec4 final = vec4(ambient + diffuse, 1) * tex;
    
    outputColor = final;
}
